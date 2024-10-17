package domain

import (
	"common/btcop"
	"common/dbutils"
	"common/op"
	"context"
	"errors"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"log"
	"time"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type WithdrawDomain struct {
	withdrawRecordRepo repo.WithdrawRecordRepo
	memberWalletDomain *MemberWalletDomain
	marketRPC          mclient.Market
	BitCoinAddress     string
}

func (d *WithdrawDomain) SaveRecord(ctx context.Context, wr *model.WithdrawRecord) error {
	return d.withdrawRecordRepo.Save(ctx, wr)
}

func (d *WithdrawDomain) Withdraw(ctx context.Context, wr model.WithdrawRecord) error {
	//1. 获取到账户的地址
	var err error
	memberWallet, err := d.memberWalletDomain.FindWalletByMemIdAndCoinId(ctx, wr.MemberId, wr.CoinId)
	if err != nil {
		return err
	}
	address := memberWallet.Address
	var txId string
	if memberWallet.CoinName == "BTC" {
		// btcTransaction完成2~6步骤
		txId, err = d.btcTransaction(address, wr.Address, wr.TotalAmount, wr.ArrivedAmount)
		if err != nil {
			return err
		}
	}
	//7. 更改数据库状态
	wr.TransactionNumber = txId
	wr.Status = 3
	wr.DealTime = time.Now().UnixMilli()
	err = d.withdrawRecordRepo.UpdateSuccess(ctx, wr)
	if err != nil {
		//TODO 如果报错，需要记录日志 要进行恢复
		log.Println(err)
	}
	return nil
}

// btcTransaction 假设 utxo = 0.1 totalAmount = 0.002 arrivedAmount = 0.00185000 fee =0.00015000
func (d *WithdrawDomain) btcTransaction(address string, toAddress string, totalAmount, arrivedAmount float64) (string, error) {
	b := &btcop.BTC{
		ApiUrl: d.BitCoinAddress,
		Auth:   "Basic Yml0Y29pbjoxMjM0NTY=",
	}
	//2. 查询账户地址的UTXO
	listUnspentResults, err := b.ListUnspent(0, 999999, []string{address})
	if err != nil {
		return "", err
	}
	//3. 判断是否符合交易
	var utxoAmount float64
	var inputs []btcop.Input
	for _, v := range listUnspentResults {
		utxoAmount += v.Amount
		inputs = append(inputs, btcop.Input{Txid: v.Txid, Vout: v.Vout})
		// 此时的unspent的数量已经达到用户的需求
		if utxoAmount >= totalAmount {
			break
		}
	}
	// 当我们遍历完全部的unspent input，仍然无法达到用户的需求，则我们的余额不足
	if utxoAmount < totalAmount {
		return "", errors.New("余额不足")
	}
	// 4. 创建交易 构建inputs和outputs
	// 计算output  两个output 假设 utxo = 0.1 totalAmount = 0.002 arrivedAmount = 0.00185000 fee =0.00015000
	// fee = allinput - alloutput = utxoAmount（0.1） - 0.00185000 - (0.1-0.002)
	var outpus []map[string]any
	oneOutput := make(map[string]any)
	oneOutput[toAddress] = arrivedAmount
	twoOutput := make(map[string]any)
	// utxoAmount = 1btc  totalAmount=0.1 ArrivedAmount=0.095 fee=0.005
	// utxoAmount-ArrivedAmount-0.005 = x
	// 1 - 0.9 -0.095 = 0.1-0.095
	twoOutput[address] = op.SubFloor(utxoAmount, totalAmount, 10)
	outpus = append(outpus, oneOutput, twoOutput)
	hexId, err := b.CreateRawTransaction(inputs, outpus)
	if err != nil {
		return "", err
	}
	//5. 签名
	sign, err := b.SignRawTransactionWithWallet(hexId)
	if err != nil {
		return "", err
	}
	//6. 发送交易到BTC网络
	txId, err := b.SendRawTransaction(sign.Hex)
	if err != nil {
		return "", err
	}
	return txId, nil
}

func (d *WithdrawDomain) RecordList(ctx context.Context, userId int64, page int64, pageSize int64) ([]*model.WithdrawRecordVo, int64, error) {
	list, total, err := d.withdrawRecordRepo.FindByUserId(ctx, userId, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var voList = make([]*model.WithdrawRecordVo, len(list))
	for i, v := range list {
		coin, err := d.marketRPC.FindCoinById(ctx, &market.MarketReq{
			Id: v.CoinId,
		})
		if err != nil {
			return nil, 0, err
		}
		voList[i] = v.ToVo(coin)
	}
	return voList, total, err
}

func NewWithdrawDomain(
	db *dbutils.ElysiaDB,
	marketRPC mclient.Market,
	BitCoinAddress string) *WithdrawDomain {
	return &WithdrawDomain{
		withdrawRecordRepo: dao.NewWithdrawRecordDao(db),
		memberWalletDomain: NewMemberWalletDomain(db),
		marketRPC:          marketRPC,
		BitCoinAddress:     BitCoinAddress,
	}
}