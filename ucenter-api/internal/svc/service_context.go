package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UCRegisterRPC ucclient.Register
	UCLoginRPC    ucclient.Login
	UCAssetRPC    ucclient.Asset
	UCMemberRPC   ucclient.Member
	UCWithdrawRPC ucclient.Withdraw
	MarketRPC     mclient.Market
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UCRegisterRPC: ucclient.NewRegister(zrpc.MustNewClient(c.UCenterRPC)),
		UCLoginRPC:    ucclient.NewLogin(zrpc.MustNewClient(c.UCenterRPC)),
		UCAssetRPC:    ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRPC)),
		UCMemberRPC:   ucclient.NewMember(zrpc.MustNewClient(c.UCenterRPC)),
		UCWithdrawRPC: ucclient.NewWithdraw(zrpc.MustNewClient(c.UCenterRPC)),
		MarketRPC:     mclient.NewMarket(zrpc.MustNewClient(c.MarketRPC)),
	}
}
