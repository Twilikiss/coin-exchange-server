package config

import (
	"exchange-api/internal/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ExchangeRPC zrpc.RpcClientConf
	Kafka       database.KafkaConfig
	JWT         AuthConfig
}

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}
