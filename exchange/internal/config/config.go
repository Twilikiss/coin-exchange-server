package config

import (
	"exchange/internal/database"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig
	CacheRedis redis.RedisConf
	Mongo      database.MongoConfig
	UCenterRPC zrpc.RpcClientConf
	MarketRPC  zrpc.RpcClientConf
	Kafka      database.KafkaConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}

type MysqlConfig struct {
	DataSource string
}
