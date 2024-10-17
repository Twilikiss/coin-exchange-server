package svc

import (
	"common/dbutils"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"market/internal/config"
	"market/internal/database"
)

type ServiceContext struct {
	Config      config.Config
	Cache       *redis.Redis
	Db          *dbutils.ElysiaDB
	MongoClient *database.MongoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := redis.MustNewRedis(c.CacheRedis)
	return &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql),
		MongoClient: database.ConnectMongo(c.Mongo),
	}
}
