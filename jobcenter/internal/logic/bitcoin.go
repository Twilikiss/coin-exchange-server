// Package logic
// @Author twilikiss 2024/8/9 0:12:12
package logic

import (
	"common/tools"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/ucclient"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"jobcenter/internal/model"
	"log"
	"strconv"
	"sync"
)

type BitCoin struct {
	wg            sync.WaitGroup
	ch            *redis.Redis
	assetRpc      ucclient.Asset
	bitCoinDomain *domain.BitCoinDomain
	queueDomain   *domain.QueueDomain
}

// Do 扫描BTC交易 查找符合系统address的交易 进行存储
func (b *BitCoin) Do(address string) {
	b.wg.Add(1)
	go b.ScanTx(address)
	b.wg.Wait()

}

func (b *BitCoin) ScanTx(btcAddress string) {
	// 1. redis是否有记录区块，获取到已处理的区块高度 dealBlocks
	dealBlocksStr, err := b.ch.Get("BTC::TX")
	if err != nil {
		logx.Error("redis 读取记录block异常， err=", err)
	}
	var dealBlocks int64
	if dealBlocksStr == "" {
		// 如果查询出来的数据为空，给与一个默认值
		dealBlocks = 2871771
	} else {
		dealBlocks = tools.ToInt64(dealBlocksStr)
	}

	// 2. 根据[getmininginfo]获取到现在的区块高度 currentBlocks
	currentBlocks, err := b.getMiningInfo(btcAddress)
	if err != nil {
		log.Println(err)
		b.wg.Done()
		return
	}
	// 3. 根据currentBlocks-dealBlocks 如果小于等于0 不需要扫描
	diff := currentBlocks - dealBlocks
	if diff <= 0 {
		b.wg.Done()
		return
	}
	// 4. 获取系统中的BTC的address列表
	ctx := context.Background()
	address, err := b.assetRpc.GetAddress(ctx, &asset.AssetReq{
		CoinName: "BTC",
	})
	if err != nil {
		log.Println(err)
		b.wg.Done()
		return
	}
	addressList := address.List
	// 5. 循环 根据getblockhash 获取 blockhash
	for i := currentBlocks; i > dealBlocks; i-- {
		blockHash, err := b.getBlockHash(i, btcAddress)
		if err != nil {
			log.Println(err)
			b.wg.Done()
			continue
		}
		// 6. 通过 getblock 获取 交易id列表
		txIdList, err := b.getBlock(blockHash, btcAddress)
		if err != nil {
			log.Println(err)
			b.wg.Done()
			continue
		}
		// 7. 循环交易id列表 获取到交易详情 得到 vout 内容
		for _, txId := range txIdList {
			txResult, err := b.getRawTransaction(txId, btcAddress)
			if err != nil {
				log.Println(err)
				b.wg.Done()
				continue
			}
			inputAddressList := make([]string, len(txResult.Vin))
			for i, vin := range txResult.Vin {
				if vin.Txid == "" {
					continue
				}
				inputTx, err := b.getRawTransaction(vin.Txid, btcAddress)
				if err != nil {
					log.Println(err)
					b.wg.Done()
					continue
				}
				vout := inputTx.Vout[vin.Vout]
				inputAddressList[i] = vout.ScriptPubKey.Address
			}
			// 8. 根据vout中的address和上方address列表进行匹配，如果匹配，我们认为是充值
			for _, vout := range txResult.Vout {
				voutAddress := vout.ScriptPubKey.Address
				flag := false
				// 9. 做一个处理，根据vint的交易 查询input的address，
				// 如果address和vout当中和系统匹配的address一样，一般认为不是充值，例如张三:2 ----> 李四:0.5 张三:1.5，这里体现出来的应该是转账（给别人充值）
				for _, inputAddress := range inputAddressList {
					if inputAddress != "" && voutAddress != "" && inputAddress == voutAddress {
						flag = true
					}
				}
				if flag {
					continue
				}
				for _, address := range addressList {
					if address != "" && address == voutAddress {
						// 数据库存有的钱包地址和区块链上的地址匹配上了
						// 充值
						// 10. 找到充值数据，存入mongo，同时发送kafka进行下一步处理（存入member_transaction表）
						err := b.bitCoinDomain.Recharge(txResult.TxId, vout.Value, voutAddress, txResult.Time, txResult.Blockhash)
						if err != nil {
							log.Println(err)
							b.wg.Done()
							continue
						}
						// 将我们的数据发送到kafka， 由其他模块进行处理
						b.queueDomain.SendRecharge(vout.Value, voutAddress, txResult.Time)
					}
				}

			}

		}
	}
	//11. 记录redis的区块高度
	err = b.ch.Set("BTC::TX", strconv.FormatInt(currentBlocks, 10))
	if err != nil {
		logx.Error("currentBlocks存入redis失败， err=", err)
	}
	b.wg.Done()
}

type MiningInfoResult struct {
	Id     string     `json:"id"`
	Error  string     `json:"error"`
	Result MiningInfo `json:"result"`
}
type MiningInfo struct {
	Blocks        int     `json:"blocks"`
	Difficulty    float64 `json:"difficulty"`
	Networkhashps float64 `json:"networkhashps"`
	Pooledtx      int     `json:"pooledtx"`
	Chain         string  `json:"chain"`
	Warnings      string  `json:"warnings"`
}

func (b *BitCoin) getMiningInfo(address string) (int64, error) {
	//{
	//    "jsonrpc": "1.0",
	//    "method": "getmininginfo",
	//    "params":[],
	//    "id": "elcoin"
	//}
	params := make(map[string]any)
	params["jsonrpc"] = "1.0"
	params["method"] = "getmininginfo"
	params["params"] = []int{}
	params["id"] = "elcoin"
	headers := make(map[string]string)
	headers["Authorization"] = "Basic Yml0Y29pbjoxMjM0NTY="
	bytes, err := tools.PostWithHeader(address, params, headers, "")
	if err != nil {
		return 0, err
	}
	var result MiningInfoResult
	_ = json.Unmarshal(bytes, &result)
	if result.Error != "" {
		return 0, errors.New(result.Error)
	}
	return int64(result.Result.Blocks), nil
}

func (b *BitCoin) getBlockHash(height int64, address string) (string, error) {
	// 这一步和getBlockInfo是类似的
	params := make(map[string]any)
	params["jsonrpc"] = "1.0"
	params["method"] = "getblockhash"
	params["params"] = []int64{height}
	params["id"] = "elcoin"
	headers := make(map[string]string)
	headers["Authorization"] = "Basic Yml0Y29pbjoxMjM0NTY="
	bytes, err := tools.PostWithHeader(address, params, headers, "")
	if err != nil {
		return "", err
	}
	var result model.BlockHashResult
	_ = json.Unmarshal(bytes, &result)
	if result.Error != "" {
		return "", errors.New(result.Error)
	}
	return result.Result, nil
}

func (b *BitCoin) getBlock(blockHash string, address string) ([]string, error) {
	params := make(map[string]any)
	params["jsonrpc"] = "1.0"
	params["method"] = "getblock"
	params["params"] = []any{blockHash, 1}
	params["id"] = "elcoin"
	headers := make(map[string]string)
	headers["Authorization"] = "Basic Yml0Y29pbjoxMjM0NTY="
	bytes, err := tools.PostWithHeader(address, params, headers, "")
	if err != nil {
		return nil, err
	}
	var result model.BlockResult
	json.Unmarshal(bytes, &result)
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return result.Result.Tx, nil
}

func (b *BitCoin) getRawTransaction(txId string, address string) (*model.RawTransaction, error) {
	params := make(map[string]any)
	params["jsonrpc"] = "1.0"
	params["method"] = "getrawtransaction"
	params["params"] = []any{txId, true}
	params["id"] = "elcoin"
	headers := make(map[string]string)
	headers["Authorization"] = "Basic Yml0Y29pbjoxMjM0NTY="
	bytes, err := tools.PostWithHeader(address, params, headers, "")
	if err != nil {
		return nil, err
	}
	var result model.RawTransactionResult
	json.Unmarshal(bytes, &result)
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result.Result, nil
}

func NewBitCoin(ch *redis.Redis, asset ucclient.Asset, mongo *database.MongoClient, kafka *database.KafkaClient) *BitCoin {
	return &BitCoin{
		ch:            ch,
		assetRpc:      asset,
		bitCoinDomain: domain.NewBitCoinDomain(mongo),
		queueDomain:   domain.NewQueueDomain(kafka),
	}
}
