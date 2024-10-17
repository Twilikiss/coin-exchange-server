package svc

import (
	"exchange-api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/exchange/eclient"
)

type ServiceContext struct {
	Config   config.Config
	OrderRPC eclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRPC: eclient.NewOrder(zrpc.MustNewClient(c.ExchangeRPC)),
	}
}
