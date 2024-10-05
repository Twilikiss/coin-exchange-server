// Package ucclient
// @Author twilikiss 2024/5/13 23:02:02
package ucclient

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"grpc-common/ucenter/types/asset"
)

type (
	AssetReq              = asset.AssetReq
	MemberWallet          = asset.MemberWallet
	MemberWalletList      = asset.MemberWalletList
	RestAddrResp          = asset.RestAddrResp
	MemberTransactionList = asset.MemberTransactionList
	AddressList           = asset.AddressList

	Asset interface {
		FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error)
		FindWallet(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWalletList, error)
		ResetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*RestAddrResp, error)
		FindAllTransaction(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberTransactionList, error)
		GetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AddressList, error)
	}

	defaultAsset struct {
		cli zrpc.Client
	}
)

func NewAsset(cli zrpc.Client) Asset {
	return &defaultAsset{
		cli: cli,
	}
}

func (m *defaultAsset) FindWalletBySymbol(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWallet, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.FindWalletBySymbol(ctx, in, opts...)
}

func (m *defaultAsset) FindWallet(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberWalletList, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.FindWallet(ctx, in, opts...)
}

func (m *defaultAsset) ResetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*RestAddrResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.ResetAddress(ctx, in, opts...)
}

func (m *defaultAsset) FindAllTransaction(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*MemberTransactionList, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.FindAllTransaction(ctx, in, opts...)
}

func (m *defaultAsset) GetAddress(ctx context.Context, in *AssetReq, opts ...grpc.CallOption) (*AddressList, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.GetAddress(ctx, in, opts...)
}
