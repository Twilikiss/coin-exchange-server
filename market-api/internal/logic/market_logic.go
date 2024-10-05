// Package logic
// @Author twilikiss 2024/5/3 0:16:16
package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	logx.Info("api register is working!")

	list = make([]*types.CoinThumbResp, 0)

	thumb := l.svcCtx.Processor.GetThumb(l.ctx)
	isCache := false
	if thumb != nil {
		switch thumb.(type) {
		case []*model.CoinThumb:

			logx.Info("缓存thumbMap已处理")

			thumbs := thumb.([]*model.CoinThumb)

			if err := copier.Copy(&list, thumbs); err != nil {
				logx.Error("copy时出现错误，error=", err)
				return nil, errors.New("内部错误，请联系管理员")
			}

			isCache = true
		}
	}

	if !isCache {
		ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
		defer cancelFunc()
		marketReq := &market.MarketReq{}
		if err := copier.Copy(marketReq, req); err != nil {
			logx.Error("copy时出现错误，error=", err)
			return nil, errors.New("内部错误，请联系管理员")
		}

		marketResp, err := l.svcCtx.MarketRPC.FindSymbolThumbTrend(ctx, marketReq)

		if err != nil {
			return nil, err
		}

		if err := copier.Copy(&list, marketResp.List); err != nil {
			logx.Error("copy时出现错误，error=", err)
			return nil, errors.New("内部错误，请联系管理员")
		}
	}
	return list, nil
}

func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	logx.Info("api register is working!")

	list = make([]*types.CoinThumbResp, 0)

	thumb := l.svcCtx.Processor.GetThumb(l.ctx)
	isCache := false
	if thumb != nil {
		switch thumb.(type) {
		case []*model.CoinThumb:

			logx.Info("缓存thumbMap已处理")

			thumbs := thumb.([]*model.CoinThumb)

			if err := copier.Copy(&list, thumbs); err != nil {
				logx.Error("copy时出现错误，error=", err)
				return nil, errors.New("内部错误，请联系管理员")
			}

			isCache = true
		}
	}

	if !isCache {
		ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
		defer cancelFunc()
		marketReq := &market.MarketReq{}
		if err := copier.Copy(marketReq, req); err != nil {
			logx.Error("copy时出现错误，error=", err)
			return nil, errors.New("内部错误，请联系管理员")
		}

		marketResp, err := l.svcCtx.MarketRPC.FindSymbolThumbTrend(ctx, marketReq)

		if err != nil {
			return nil, err
		}

		if err := copier.Copy(&list, marketResp.List); err != nil {
			logx.Error("copy时出现错误，error=", err)
			return nil, errors.New("内部错误，请联系管理员")
		}
	}
	return list, nil
}

func (l *MarketLogic) SymbolInfo(req *types.MarketReq) (*types.ExchangeCoinResp, error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()
	marketReq := &market.MarketReq{
		Ip:     req.Ip,
		Symbol: req.Symbol,
	}
	if err := copier.Copy(marketReq, req); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	resp, err := l.svcCtx.MarketRPC.FindSymbolInfo(ctx, marketReq)

	if err != nil {
		return nil, err
	}

	exchangeCoinResp := &types.ExchangeCoinResp{}

	if err := copier.Copy(exchangeCoinResp, resp); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	return exchangeCoinResp, nil
}

func (l *MarketLogic) CoinInfo(req *types.MarketReq) (*types.Coin, error) {
	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()
	marketReq := &market.MarketReq{
		Unit: req.Unit,
	}
	resp, err := l.svcCtx.MarketRPC.FindCoinInfo(ctx, marketReq)

	if err != nil {
		return nil, err
	}

	coin := &types.Coin{}

	if err := copier.Copy(coin, resp); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	return coin, nil
}

func (l *MarketLogic) History(req *types.MarketReq) (*types.HistoryKline, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	historyKline, err := l.svcCtx.MarketRPC.HistoryKline(ctx, &market.MarketReq{
		Symbol:     req.Symbol,
		From:       req.From,
		To:         req.To,
		Resolution: req.Resolution,
	})
	if err != nil {
		return nil, err
	}
	histories := historyKline.List
	var list = make([][]any, len(histories))
	for i, v := range histories {
		content := make([]any, 6)
		content[0] = v.Time
		content[1] = v.Open
		content[2] = v.High
		content[3] = v.Low
		content[4] = v.Close
		content[5] = v.Volume
		list[i] = content
	}
	data := &types.HistoryKline{}
	*data = list
	return data, nil
}
