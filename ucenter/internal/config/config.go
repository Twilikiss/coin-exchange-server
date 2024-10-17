package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"ucenter/internal/database"
	"ucenter/internal/model"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql       MysqlConfig
	CacheRedis  redis.RedisConf
	Captcha     CaptchaConf
	JWT         AuthConfig
	MarketRPC   zrpc.RpcClientConf
	Kafka       database.KafkaConfig
	ExchangeRPC zrpc.RpcClientConf
	Bitcoin     model.BitCoinConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}

type MysqlConfig struct {
	DataSource string
}
type CaptchaConf struct {
	Vid       string
	SecretKey string
}
