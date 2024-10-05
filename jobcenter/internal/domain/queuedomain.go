// Package domain
// @Author twilikiss 2024/5/8 17:19:19
package domain

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
)

const KLINE1M = "kline_1m"
const BtcTransactionTopic = "BTC_TRANSACTION"

type QueueDomain struct {
	kafkaCli *database.KafkaClient
}

func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	kline := model.NewKline(data, "1m")
	bytes, _ := json.Marshal(kline)
	msg := database.KafkaData{
		Topic: KLINE1M,
		Key:   []byte(symbol),
		Data:  bytes,
	}
	d.kafkaCli.Send(msg)
	logx.Info("================数据已发送至kafka=============")
}

func (d *QueueDomain) SendRecharge(value float64, address string, time int64) {
	data := make(map[string]any)
	data["value"] = value
	data["address"] = address
	data["time"] = time
	data["type"] = model.RECHARGE
	data["symbol"] = "BTC"
	marshal, _ := json.Marshal(data)
	msg := database.KafkaData{
		Topic: BtcTransactionTopic,
		Data:  marshal,
		Key:   []byte(address),
	}
	d.kafkaCli.Send(msg)
	logx.Info("================Transaction数据已发送至kafka=============")
}

func NewQueueDomain(kafkaCli *database.KafkaClient) *QueueDomain {
	return &QueueDomain{
		kafkaCli: kafkaCli,
	}
}
