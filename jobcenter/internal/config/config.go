package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
)

type Config struct {
	Okx        model.OkxConfig
	Mongo      database.MongoConfig
	Kafka      database.KafkaConfig
	CacheRedis redis.RedisConf
	UCenterRPC zrpc.RpcClientConf
	Bitcoin    model.BitCoinConfig
}
