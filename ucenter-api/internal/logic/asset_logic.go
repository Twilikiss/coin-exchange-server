// Package logic
// @Author twilikiss 2024/5/13 22:48:48
package logic

import (
	"common/pages"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/asset"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type AssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssetLogic) FindWalletBySymbol(req *types.AssetReq) (*types.MemberWallet, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	memberWallet, err := l.svcCtx.UCAssetRPC.FindWalletBySymbol(ctx, &asset.AssetReq{
		CoinName: req.CoinName,
		UserId:   value,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.MemberWallet{}
	if err := copier.Copy(resp, memberWallet); err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *AssetLogic) FindWallet(req *types.AssetReq) ([]*types.MemberWallet, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	memberWallets, err := l.svcCtx.UCAssetRPC.FindWallet(ctx, &asset.AssetReq{
		UserId: value,
	})
	if err != nil {
		return nil, err
	}
	var resp []*types.MemberWallet
	if err := copier.Copy(&resp, memberWallets.List); err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *AssetLogic) RestAddress(req *types.AssetReq) (string, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	_, err := l.svcCtx.UCAssetRPC.ResetAddress(ctx, &asset.AssetReq{
		UserId:   value,
		CoinName: req.Unit,
	})
	if err != nil {
		return "", err
	}
	return "", nil
}

func (l *AssetLogic) FindAllTransaction(req *types.AssetReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	resp, err := l.svcCtx.UCAssetRPC.FindAllTransaction(ctx, &asset.AssetReq{
		UserId:    value,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageNo:    int64(req.PageNo),
		PageSize:  int64(req.PageSize),
		Type:      req.Type,
		Symbol:    req.Symbol,
	})
	if err != nil {
		return nil, err
	}

	total := resp.Total
	b := make([]any, len(resp.List))
	for i, v := range resp.List {
		b[i] = v
	}

	return pages.New(b, int64(req.PageNo), int64(req.PageSize), total), nil
}
