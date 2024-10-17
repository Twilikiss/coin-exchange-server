// Package server
// @Author twilikiss 2024/5/13 23:08:08
package server

import (
	"context"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type AssetServer struct {
	svcCtx *svc.ServiceContext
	asset.UnimplementedAssetServer
}

func NewAssetServer(svcCtx *svc.ServiceContext) *AssetServer {
	return &AssetServer{
		svcCtx: svcCtx,
	}
}

func (s *AssetServer) FindWalletBySymbol(ctx context.Context, in *asset.AssetReq) (*asset.MemberWallet, error) {
	l := logic.NewAssetLogic(ctx, s.svcCtx)
	return l.FindWalletBySymbol(in)
}

func (s *AssetServer) FindWallet(ctx context.Context, in *asset.AssetReq) (*asset.MemberWalletList, error) {
	l := logic.NewAssetLogic(ctx, s.svcCtx)
	return l.FindWallet(in)
}

func (s *AssetServer) ResetAddress(ctx context.Context, in *asset.AssetReq) (*asset.RestAddrResp, error) {
	l := logic.NewAssetLogic(ctx, s.svcCtx)
	return l.ResetAddress(in)
}
func (s *AssetServer) FindAllTransaction(ctx context.Context, in *asset.AssetReq) (*asset.MemberTransactionList, error) {
	l := logic.NewAssetLogic(ctx, s.svcCtx)
	return l.FindAllTransaction(in)
}
func (s *AssetServer) GetAddress(ctx context.Context, in *asset.AssetReq) (*asset.AddressList, error) {
	l := logic.NewAssetLogic(ctx, s.svcCtx)
	return l.GetAddress(in)
}
