package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"market/internal/database"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      database.MysqlConfig
	CacheRedis redis.RedisConf
	Mongo      database.MongoConfig
}
