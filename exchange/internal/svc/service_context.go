package svc

import (
	"common/dbutils"
	"exchange/internal/config"
	"exchange/internal/consumer"
	"exchange/internal/database"
	"exchange/internal/processor"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/ucclient"
)

type ServiceContext struct {
	Config      config.Config
	Cache       *redis.Redis
	Db          *dbutils.ElysiaDB
	MongoClient *database.MongoClient
	MemberRPC   ucclient.Member
	MarketRPC   mclient.Market
	AssetRPC    ucclient.Asset
	KafkaClient *database.KafkaClient
}

// init 设置当前服务端自动调用RPC或者其他MQ的操作
func (s *ServiceContext) init() {
	factory := processor.NewCoinTradeFactory()
	factory.Init(s.MarketRPC, s.KafkaClient, s.Db)
	kafkaConsumer := consumer.NewKafkaConsumer(s.KafkaClient, factory, s.Db)
	kafkaConsumer.Run()
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := redis.MustNewRedis(c.CacheRedis)
	kafkaClient := database.NewKafkaClient(c.Kafka)
	s := &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql.DataSource),
		MongoClient: database.ConnectMongo(c.Mongo),
		MemberRPC:   ucclient.NewMember(zrpc.MustNewClient(c.UCenterRPC)),
		MarketRPC:   mclient.NewMarket(zrpc.MustNewClient(c.MarketRPC)),
		AssetRPC:    ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRPC)),
		KafkaClient: kafkaClient,
	}
	s.init()
	return s
}
