// Package consumer
// @Author twilikiss 2024/6/25 16:48:48
package consumer

import (
	"common/dbutils"
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/domain"
	"exchange/internal/model"
	"exchange/internal/processor"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type KafkaConsumer struct {
	cli     *database.KafkaClient
	factory *processor.CoinTradeFactory
	db      *dbutils.ElysiaDB
}

func NewKafkaConsumer(cli *database.KafkaClient, factory *processor.CoinTradeFactory, db *dbutils.ElysiaDB) *KafkaConsumer {
	return &KafkaConsumer{
		cli:     cli,
		factory: factory,
		db:      db,
	}
}

// 消费订单消息，拿到我们完成创建操作的订单
// 1. 先实现买卖盘的逻辑 买 卖 一旦匹配完成 成交了 成交的价格和数量  就会做为别人的参考 买卖盘也是实时

// Run 对外暴露的启动方法
func (k *KafkaConsumer) Run() {
	exchangeOrderDomain := domain.NewExchangeOrderDomain(k.db)
	k.orderTrading()
	k.orderComplete(exchangeOrderDomain)
}

func (k *KafkaConsumer) orderTrading() {
	kafkaClient := k.cli.StartRead("exchange_order_trading")
	// 开启循环读取kafka中的数据
	go k.readOrderTrading(kafkaClient)
}

func (k *KafkaConsumer) readOrderTrading(client *database.KafkaClient) {
	for {
		kafkadata := client.Read()
		var order *model.ExchangeOrder
		err := json.Unmarshal(kafkadata.Data, &order)
		if err != nil {
			logx.Error("json解码转换失败，err=", err)
			continue
		}

		// ===================================将数据交给撮合交易引擎处理=======================================
		// 得到撮合交易引擎
		coinTrade := k.factory.GetCoinTrade(order.Symbol)
		coinTrade.Trade(order)
	}
}

func (k *KafkaConsumer) orderComplete(orderDomain *domain.ExchangeOrderDomain) {
	cli := k.cli.StartRead("exchange_order_complete")
	go k.readOrderComplete(cli, orderDomain)
}

func (k *KafkaConsumer) readOrderComplete(cli *database.KafkaClient, orderDomain *domain.ExchangeOrderDomain) {
	for {
		kafkaData := cli.Read()
		var order *model.ExchangeOrder
		err := json.Unmarshal(kafkaData.Data, &order)
		if err != nil {
			logx.Error("json解码错误，err=", err)
		}
		//这个时候 我们需要去更改状态
		err = orderDomain.UpdateOrderComplete(context.Background(), order)
		if err != nil {
			logx.Error("订单状态修改失败，err=", err)
			cli.RPut(kafkaData)
			time.Sleep(200 * time.Millisecond)
			continue
		}

		// 通过kafka通知钱包更新
		for {
			kafkaData.Topic = "exchange_order_complete_update_success"
			err2 := cli.SendSync(kafkaData)
			if err2 != nil {
				logx.Error(err2)
				time.Sleep(250 * time.Millisecond)
				continue
			}
			break
		}
	}
}
