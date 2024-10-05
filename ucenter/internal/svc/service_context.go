package svc

import (
	"common/dbutils"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/exchange/eclient"
	"grpc-common/market/mclient"
	"ucenter/internal/config"
	"ucenter/internal/consumer"
	"ucenter/internal/database"
)

type ServiceContext struct {
	Config      config.Config
	Cache       *redis.Redis
	Db          *dbutils.ElysiaDB
	MarketRPC   mclient.Market
	KafkaClient *database.KafkaClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := redis.MustNewRedis(c.CacheRedis)
	mysql := database.ConnMysql(c.Mysql.DataSource)
	cli := database.NewKafkaClient(c.Kafka)
	orderCli := cli.StartReadNew("add-exchange-order")
	order := eclient.NewOrder(zrpc.MustNewClient(c.ExchangeRPC))
	go consumer.ExchangeOrderAdd(redisCache, orderCli, order, mysql)

	completeCli := cli.StartReadNew("exchange_order_complete_update_success")
	go consumer.ExchangeOrderComplete(redisCache, completeCli, mysql)

	btCli := cli.StartReadNew("BTC_TRANSACTION")
	go consumer.BitCoinTransaction(redisCache, btCli, mysql)

	withdrawCli := cli.StartReadNew("withdraw")
	go consumer.WithdrawConsumer(withdrawCli, mysql, c.Bitcoin.Address)

	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite() // 注意要开启写操作
	return &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql.DataSource),
		MarketRPC:   mclient.NewMarket(zrpc.MustNewClient(c.MarketRPC)),
		KafkaClient: client,
	}
}
