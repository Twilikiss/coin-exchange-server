// Package domain
// @Author twilikiss 2024/5/19 12:07:07
package domain

import (
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type KafkaDomain struct {
	cli         *database.KafkaClient
	orderDomain *ExchangeOrderDomain
}

func NewKafkaDomain(cli *database.KafkaClient, domain *ExchangeOrderDomain) *KafkaDomain {
	k := &KafkaDomain{
		cli:         cli,
		orderDomain: domain,
	}
	go k.WaitAddOrderResult()
	return k
}

// SendOrderAdd 发送创建的订单的相关消息到kafka
func (d *KafkaDomain) SendOrderAdd(
	topic string,
	userId int64,
	orderId string,
	money float64,
	symbol string,
	direction int,
	baseSymbol string,
	coinSymbol string) error {

	m := make(map[string]any)
	m["userId"] = userId
	m["orderId"] = orderId
	m["money"] = money
	m["symbol"] = symbol
	m["direction"] = direction
	m["baseSymbol"] = baseSymbol
	m["coinSymbol"] = coinSymbol
	marshal, _ := json.Marshal(m)
	data := database.KafkaData{
		Topic: topic,
		Key:   []byte(orderId),
		Data:  marshal,
	}
	err := d.cli.SendSync(data)
	logx.Info("创建订单，发消息成功,orderId=" + orderId)
	return err
}

type OrderResult struct {
	UserId  int64  `json:"userId"`
	OrderId string `json:"orderId"`
}

// WaitAddOrderResult 等待用户创建订单后，冻结相关钱包后，再将我们的订单状态修改为【可交易】
func (d *KafkaDomain) WaitAddOrderResult() {
	// 开启kafka读取
	kafkaClient := d.cli.StartRead("exchange_order_init_complete_trading")
	for {
		kafkaData := kafkaClient.Read()
		logx.Info("读取exchange_order_init_complete_trading消息成功：", string(kafkaData.Key))
		var orderResult OrderResult
		err := json.Unmarshal(kafkaData.Data, &orderResult)
		if err != nil {
			logx.Error("转换错误， err=", err)
			continue
		}
		exchangeOrder, err := d.orderDomain.FindOrderByOrderId(context.Background(), orderResult.OrderId)

		// =============================对参数进行校验==================================

		if err != nil {
			logx.Error(err)
			// 如果出现错误的话，将我们的Order状态切换为【取消交易】
			err := d.orderDomain.UpdateStatusCancel(context.Background(), orderResult.OrderId)
			if err != nil {
				logx.Error("修改Order状态【取消交易】错误，err=", err)
				// 重新放入到MQ中进行消费
				kafkaClient.RPut(kafkaData)
				time.Sleep(500 * time.Millisecond)
			}
			continue
		}

		if exchangeOrder == nil {
			logx.Error("订单id不存在, id=", orderResult.OrderId)
			continue
		}

		if exchangeOrder.Status != model.Init {
			logx.Error("订单已经被处理过了")
			continue
		}

		// 将Order状态切换为【可交易】
		err = d.orderDomain.UpdateStatusTrading(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error("修改Order状态【可交易】错误，err=", err)
			// 重新放入到MQ中进行消费
			kafkaClient.RPut(kafkaData)
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// 改变对应状态
		exchangeOrder.Status = model.Trading

		//订单初始化完成 发送消息到kafka 等待撮合交易引擎进行交易撮合
		for {
			bytes, _ := json.Marshal(exchangeOrder)
			orderData := database.KafkaData{
				Topic: "exchange_order_trading",
				Key:   []byte(exchangeOrder.OrderId),
				Data:  bytes,
			}
			err := d.cli.SendSync(orderData)
			if err != nil {
				logx.Error("撮合交易相关数据发送失败，err=", err)
				time.Sleep(250 * time.Millisecond)
				continue
			}
			logx.Info("撮合交易相关数据发送成功，发送成功消息:", exchangeOrder.OrderId)
			break
		}
	}
}
