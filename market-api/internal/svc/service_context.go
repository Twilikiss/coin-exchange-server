package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"market-api/internal/config"
	"market-api/internal/database"
	"market-api/internal/processor"
	"market-api/internal/ws"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRPC mclient.ExchangeRate
	MarketRPC       mclient.Market
	Processor       processor.Processor
}

func NewServiceContext(c config.Config, wsServer *ws.WebsocketServer) *ServiceContext {
	// 初始化processor
	kafkaCli := database.NewKafkaClient(c.Kafka)
	market := mclient.NewMarket(zrpc.MustNewClient(c.MarketRPC))
	defaultProcessor := processor.NewDefaultProcessor(market, kafkaCli)
	defaultProcessor.Init()
	defaultProcessor.AddHandler(processor.NewWebSocketHandler(wsServer))
	return &ServiceContext{
		Config:          c,
		ExchangeRateRPC: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRPC)),
		MarketRPC:       market,
		Processor:       defaultProcessor,
	}
}
