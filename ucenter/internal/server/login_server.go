// Package server
// @Author twilikiss 2024/5/3 1:18:18
package server

import (
	"context"
	"grpc-common/ucenter/types/login"
	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	login.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (ls *LoginServer) Login(ctx context.Context, in *login.LoginReq) (*login.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, ls.svcCtx)
	return l.Login(in)
}
