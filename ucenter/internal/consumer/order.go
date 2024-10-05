// Package consumer
// @Author twilikiss 2024/5/20 16:48:48
package consumer

import (
	"common/dbutils"
	"common/dbutils/tran"
	"common/enum"
	"common/op"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"grpc-common/exchange/eclient"
	"grpc-common/exchange/types/order"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
)

type OrderAdd struct {
	UserId     int64   `json:"userId"`
	OrderId    string  `json:"orderId"`
	Money      float64 `json:"money"`
	Symbol     string  `json:"symbol"`
	Direction  int     `json:"direction"`
	BaseSymbol string  `json:"baseSymbol"`
	CoinSymbol string  `json:"coinSymbol"`
}

// ExchangeOrderAdd 订单添加
// redisCli *redis.Redis 主要为Redis提供
func ExchangeOrderAdd(redisCli *redis.Redis, cli *database.KafkaClient, orderRpc eclient.Order, db *dbutils.ElysiaDB) {
	for {
		// =========================获取数据并解码得到原始数据=============================
		kafkaData := cli.Read()
		orderId := string(kafkaData.Key)
		logx.Info("ucenter接收到Kafka数据，正在对用户的钱包进行处理，orderId=", orderId)
		var orderAdd OrderAdd
		err := json.Unmarshal(kafkaData.Data, &orderAdd)
		if err != nil {
			logx.Error("转换数据出错，err=")
			continue
		}
		if orderId != orderAdd.OrderId {
			logx.Error("消息数据有误")
			continue
		}

		// ==========================获取数据并解码得到原始数据 END========================== //

		// 冻结用户的相关资金
		ctx := context.Background()
		exchangeOrderOrigin, err := orderRpc.FindByOrderId(ctx, &order.OrderReq{OrderId: orderId})
		if err != nil {
			logx.Error("failed to orderRpc, err=", err)
			// 如果出现错误，显然用户钱包冻结失败
			cancelOrder(ctx, orderId, orderRpc, kafkaData, cli)
			continue
		}
		if exchangeOrderOrigin == nil {
			logx.Error("orderId = " + orderId + "不存在")
			continue
		}
		// Init状态:4
		if exchangeOrderOrigin.Status != 4 {
			logx.Error("orderId = " + orderId + "已经被操作过了")
			continue
		}

		// 利用Redis来实现简单高效的分布式锁功能
		lock := redis.NewRedisLock(redisCli, "exchange_order::"+fmt.Sprintf("%d::%s", orderAdd.UserId, orderId))
		//查询订单信息 如果是正在交易中 继续 否则return
		acquireCtx, err := lock.AcquireCtx(ctx)
		if err != nil {
			logx.Error(err)
			logx.Info("已经有别的进程处理此消息")
			continue
		}

		// 获取分布式锁
		if acquireCtx {

			// 因为涉及到用户金额的相关操作，因此要引入事务的相关操作
			transaction := tran.NewTransaction(db.Conn)
			err = transaction.Action(func(conn dbutils.DbConn) error {
				// BUY = 0
				walletDomain := domain.NewMemberWalletDomain(db)
				if orderAdd.Direction == 0 {
					err = walletDomain.Freeze(ctx, conn, orderAdd.UserId, orderAdd.Money, orderAdd.BaseSymbol)
				} else { // 如果是卖的话
					err = walletDomain.Freeze(ctx, conn, orderAdd.UserId, orderAdd.Money, orderAdd.CoinSymbol)
				}
				if err != nil {
					// 如果出现错误，显然用户钱包冻结失败
					return err
				}
				return nil
			})

			if err != nil {
				// 如果出现错误
				cancelOrder(ctx, orderId, orderRpc, kafkaData, cli)
				continue
			}

			logx.Infof("orderId=%s，钱包处理完成", orderId)

			//需要将状态 改为trading
			//都完成后 通知订单进行状态变更 需要保证一定发送成功
			for {
				m := make(map[string]any)
				m["userId"] = orderAdd.UserId
				m["orderId"] = orderId
				marshal, _ := json.Marshal(m)
				data := database.KafkaData{
					Topic: "exchange_order_init_complete_trading",
					Key:   []byte(orderId),
					Data:  marshal,
				}
				err := cli.SendSync(data)
				if err != nil {
					logx.Error(err)
					time.Sleep(250 * time.Millisecond)
					continue
				}
				logx.Info("发送exchange_order_init_complete_trading 消息成功:" + orderId)
				break
			}
			_, err := lock.Release()
			if err != nil {
				logx.Error("lock释放出错，err=", err)
			}
		}
	}
}

// TODO 这里最好传一个order的状态，需要校验状态一致才能进行cancel操作（？）
func cancelOrder(ctx context.Context, orderId string, orderRpc eclient.Order, kafkaData database.KafkaData, cli *database.KafkaClient) {
	_, err := orderRpc.CancelOrder(ctx, &order.OrderReq{
		OrderId: orderId,
	})
	if err != nil {
		// 如果连我们的回滚都失败，只能将数据重新放回到kafka中重新消费
		cli.RPut(kafkaData)
	}
}

type ExchangeOrder struct {
	Id            int64   `gorm:"column:id" json:"id"`
	OrderId       string  `gorm:"column:order_id" json:"orderId"`
	Amount        float64 `gorm:"column:amount" json:"amount"`
	BaseSymbol    string  `gorm:"column:base_symbol" json:"baseSymbol"`
	CanceledTime  int64   `gorm:"column:canceled_time" json:"canceledTime"`
	CoinSymbol    string  `gorm:"column:coin_symbol" json:"coinSymbol"`
	CompletedTime int64   `gorm:"column:completed_time" json:"completedTime"`
	Direction     int     `gorm:"column:direction" json:"direction"`
	MemberId      int64   `gorm:"column:member_id" json:"memberId"`
	Price         float64 `gorm:"column:price" json:"price"`
	Status        int     `gorm:"column:status" json:"status"`
	Symbol        string  `gorm:"column:symbol" json:"symbol"`
	Time          int64   `gorm:"column:time" json:"time"`
	TradedAmount  float64 `gorm:"column:traded_amount" json:"tradedAmount"`
	Turnover      float64 `gorm:"column:turnover" json:"turnover"`
	Type          int     `gorm:"column:type" json:"type"`
	UseDiscount   string  `gorm:"column:use_discount" json:"useDiscount"`
}

// status
const (
	Trading = iota
	Completed
	Canceled
	OverTimed
	Init
)

var StatusMap = enum.Enum{
	Trading:   "TRADING",
	Completed: "COMPLETED",
	Canceled:  "CANCELED",
	OverTimed: "OVERTIMED",
}

// direction
const (
	BUY = iota
	SELL
)

var DirectionMap = enum.Enum{
	BUY:  "BUY",
	SELL: "SELL",
}

// type
const (
	MarketPrice = iota
	LimitPrice
)

var TypeMap = enum.Enum{
	MarketPrice: "MARKET_PRICE",
	LimitPrice:  "LIMIT_PRICE",
}

func ExchangeOrderComplete(redisCli *redis.Redis, cli *database.KafkaClient, db *dbutils.ElysiaDB) {
	//先接收消息
	for {
		kafkaData := cli.Read()
		var exchangeOrder *ExchangeOrder
		err := json.Unmarshal(kafkaData.Data, &exchangeOrder)
		if err != nil {
			logx.Error("json解码失败，err=", err)
		}
		if exchangeOrder == nil {
			continue
		}
		if exchangeOrder.Status != Completed {
			continue
		}
		logx.Info("收到exchange_order_complete_update_success 消息成功:" + exchangeOrder.OrderId)
		walletDomain := domain.NewMemberWalletDomain(db)
		lock := redis.NewRedisLock(redisCli, fmt.Sprintf("order_complete_update_wallet::%d", exchangeOrder.MemberId))
		acquire, err := lock.Acquire()
		if err != nil {
			logx.Error(err)
			logx.Info("有进程已经拿到锁进行处理了")
			continue
		}

		if acquire {
			ctx := context.Background()
			if exchangeOrder.Direction == BUY {

				baseWallet, err := walletDomain.FindWalletByMemIdAndCoin(ctx, exchangeOrder.MemberId, exchangeOrder.BaseSymbol)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}

				coinWallet, err := walletDomain.FindWalletByMemIdAndCoin(ctx, exchangeOrder.MemberId, exchangeOrder.CoinSymbol)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}

				if exchangeOrder.Type == MarketPrice {
					//市价买 amount USDT冻结的钱   order.turnover扣的钱  amount-order.turnover还回去的钱
					// 解除已成交部分的冻结
					baseWallet.FrozenBalance = op.SubFloor(baseWallet.FrozenBalance, exchangeOrder.Amount, 8)
					// 把未交易的金额再加回去
					baseWallet.Balance = op.AddFloor(baseWallet.Balance, op.SubFloor(exchangeOrder.Amount, exchangeOrder.Turnover, 8), 8)
					// 把买来的部分加入到我们对应的钱包
					coinWallet.Balance = op.AddFloor(coinWallet.Balance, exchangeOrder.TradedAmount, 8)
				} else {
					//限价买 冻结的钱是 order.price*amount  成交了turnover 还回去的钱 order.price*amount-order.turnover
					floor := op.MulFloor(exchangeOrder.Price, exchangeOrder.Amount, 8)
					baseWallet.FrozenBalance = op.SubFloor(baseWallet.FrozenBalance, floor, 8)
					baseWallet.Balance = op.AddFloor(baseWallet.Balance, op.SubFloor(floor, exchangeOrder.Turnover, 8), 8)
					coinWallet.Balance = op.AddFloor(coinWallet.Balance, exchangeOrder.TradedAmount, 8)
				}
				err = walletDomain.UpdateWalletCoinAndBase(ctx, baseWallet, coinWallet)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}
			} else {
				// 卖相关的订单
				// 卖 不管是市价还是限价 都是卖的 BTC  解冻amount 得到的钱是 order.turnover
				coinWallet, err := walletDomain.FindWalletByMemIdAndCoin(ctx, exchangeOrder.MemberId, exchangeOrder.CoinSymbol)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}
				baseWallet, err := walletDomain.FindWalletByMemIdAndCoin(ctx, exchangeOrder.MemberId, exchangeOrder.BaseSymbol)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}

				coinWallet.FrozenBalance = op.SubFloor(coinWallet.FrozenBalance, exchangeOrder.Amount, 8)
				baseWallet.Balance = op.AddFloor(baseWallet.Balance, exchangeOrder.Turnover, 8)
				err = walletDomain.UpdateWalletCoinAndBase(ctx, baseWallet, coinWallet)
				if err != nil {
					logx.Error(err)
					cli.RPut(kafkaData)
					time.Sleep(250 * time.Millisecond)
					lock.Release()
					continue
				}
			}
			logx.Info("更新钱包成功:" + exchangeOrder.OrderId)
			lock.Release()
		}
	}
}
