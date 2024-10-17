// Code generated by goctl. DO NOT EDIT.
// Source: register.proto

package ucclient

import (
	"context"
	"grpc-common/ucenter/types/login"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginReq = login.LoginReq
	LoginRes = login.LoginRes

	Login interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	}

	defaultLogin struct {
		cli zrpc.Client
	}
)

func NewLogin(cli zrpc.Client) Login {
	return &defaultLogin{
		cli: cli,
	}
}

func (m *defaultLogin) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	client := login.NewLoginClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}