// Package svc
// @Author twilikiss 2024/5/5 12:40:40
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/ucclient"
	"jobcenter/internal/config"
	"jobcenter/internal/database"
)

type ServiceContext struct {
	Config         config.Config
	MongoClient    *database.MongoClient
	KafkaClient    *database.KafkaClient
	Cache          *redis.Redis
	AssetRpc       ucclient.Asset
	BitCoinAddress string
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite() // 注意要开启写操作
	redisCache := redis.MustNewRedis(c.CacheRedis)
	return &ServiceContext{
		Config:         c,
		MongoClient:    database.ConnectMongo(c.Mongo),
		KafkaClient:    client,
		Cache:          redisCache,
		AssetRpc:       ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRPC)),
		BitCoinAddress: c.Bitcoin.Address,
	}
}
