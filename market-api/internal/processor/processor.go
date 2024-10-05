// Package processor
// @Author twilikiss 2024/5/8 20:08:08
package processor

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"
	"sync"
	"time"
)

const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"
const TradePlateTopic = "exchange_order_trade_plate"
const TradePlate = "tradePlate"

type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline, thumb map[string]*model.CoinThumb)
	HandleTradePlate(symbol string, tp *model.TradePlateResult)
}

type Processor interface {
	GetThumb(ctx context.Context) any
	RefreshThumbMap()
	Process(data ProcessData)
	AddHandler(h MarketHandler) // 处理Kline数据
}

type DefaultProcessor struct {
	lock      sync.RWMutex
	kafkaCli  *database.KafkaClient
	handlers  []MarketHandler
	thumbMap  map[string]*model.CoinThumb
	marketRpc mclient.Market
}

func NewDefaultProcessor(marketRpc mclient.Market, kafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		kafkaCli:  kafkaCli,
		handlers:  make([]MarketHandler, 0),
		thumbMap:  make(map[string]*model.CoinThumb),
		marketRpc: marketRpc,
	}
}

func (d *DefaultProcessor) Init() {
	d.startReadFromKafka(KLINE1M, KLINE)
	d.startReadTradePlate(TradePlateTopic)
	d.initThumbMap()
}

func (d *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	// 要注意我们需要startRead后才能通过read获取数据
	d.kafkaCli.StartRead(topic) // 这里我们补全我们的topic信息
	// 处理读取后的数据
	go d.dealQueueData(d.kafkaCli, tp)
}

func (d *DefaultProcessor) startReadTradePlate(topic string) {
	// 要注意我们需要startRead后才能通过read获取数据
	cli := d.kafkaCli.StartReadNew(topic) // 这里我们补全我们的topic信息
	// 处理读取后的数据
	go d.dealQueueData(cli, TradePlate)
}

func (d *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	// 读取我们队列的数据
	for {
		// 当我们的kafka中接收到数据时就要触发websocket
		msg := cli.Read()
		data := ProcessData{
			Type: tp,
			Key:  msg.Key,
			Data: msg.Data,
		}
		d.Process(data)
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	d.handlers = append(d.handlers, h)
}

func (d *DefaultProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		_ = json.Unmarshal(data.Data, kline)
		// 调用MarkHandler来处理我们的数据
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumbMap)
		}
	} else if data.Type == TradePlate {
		symbol := string(data.Key)
		tp := &model.TradePlateResult{}
		_ = json.Unmarshal(data.Data, tp)
		for _, v := range d.handlers {
			v.HandleTradePlate(symbol, tp)
		}
	}
}

func (d *DefaultProcessor) GetThumb(ctx context.Context) any {
	d.lock.Lock()
	d.RefreshThumbMap()
	cs := make([]*model.CoinThumb, len(d.thumbMap))
	i := 0
	for _, v := range d.thumbMap {
		cs[i] = v
		i++
	}
	d.lock.Unlock()
	return cs
}

func (d *DefaultProcessor) RefreshThumbMap() {
	logx.Info("ThumbMap已更新并缓存啦")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	symbolThumbRes, err := d.marketRpc.FindSymbolThumbTrend(
		ctx,
		&market.MarketReq{Period: "1m"})
	if err != nil {
		logx.Info(err)
	} else {
		coinThumbs := symbolThumbRes.List
		for _, v := range coinThumbs {
			data := &model.CoinThumb{}
			err := copier.Copy(data, v)
			if err != nil {
				logx.Error("转换失败，err=", err)
				return
			}
			d.thumbMap[v.Symbol] = data
		}
	}
}

func (d *DefaultProcessor) initThumbMap() {
	logx.Info("############## initThumbMap is running! ##################")
	d.RefreshThumbMap()
}
