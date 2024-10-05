// Package ucclient
// @Author twilikiss 2024/8/14 0:06:06
package ucclient

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"grpc-common/ucenter/types/withdraw"
)

type (
	WithdrawReq       = withdraw.WithdrawReq
	AddressSimpleList = withdraw.AddressSimpleList
	WithdrawNoRes     = withdraw.NoRes
	RecordList        = withdraw.RecordList

	Withdraw interface {
		FindAddressByCoinId(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*AddressSimpleList, error)
		SendCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawNoRes, error)
		WithdrawCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawNoRes, error)
		WithdrawRecord(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*RecordList, error)
	}

	defaultWithdraw struct {
		cli zrpc.Client
	}
)

func (d *defaultWithdraw) WithdrawRecord(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*RecordList, error) {
	client := withdraw.NewWithdrawClient(d.cli.Conn())
	return client.WithdrawRecord(ctx, in, opts...)
}

func (d *defaultWithdraw) WithdrawCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawNoRes, error) {
	client := withdraw.NewWithdrawClient(d.cli.Conn())
	return client.WithdrawCode(ctx, in, opts...)
}

func (d *defaultWithdraw) SendCode(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*WithdrawNoRes, error) {
	client := withdraw.NewWithdrawClient(d.cli.Conn())
	return client.SendCode(ctx, in, opts...)
}

func (d *defaultWithdraw) FindAddressByCoinId(ctx context.Context, in *WithdrawReq, opts ...grpc.CallOption) (*AddressSimpleList, error) {
	client := withdraw.NewWithdrawClient(d.cli.Conn())
	return client.FindAddressByCoinId(ctx, in, opts...)
}

func NewWithdraw(cli zrpc.Client) Withdraw {
	return &defaultWithdraw{
		cli: cli,
	}
}
