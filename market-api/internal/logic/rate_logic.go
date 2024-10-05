// Package logic
// @Author twilikiss 2024/5/3 0:16:16
package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/rate"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

type ExchangeRateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExchangeRateLogic) UsdRate(req *types.RateRequest) (resp *types.RateResponse, err error) {
	logx.Info("api register is working!")
	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()

	rateReq := &rate.RateReq{}
	if err := copier.Copy(rateReq, req); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	rateResp, err := l.svcCtx.ExchangeRateRPC.UsdRate(ctx, rateReq)

	if err != nil {
		return nil, err
	}

	data := &types.RateResponse{}

	if err := copier.Copy(data, rateResp); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	return data, nil
}
