// Package domain
// @Author twilikiss 2024/5/7 19:53:53
package domain

import (
	"common/op"
	"common/tools"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"market/internal/dao"
	"market/internal/database"
	"market/internal/model"
	"market/internal/repo"
	"time"
)

type MarketDomain struct {
	klineRepo repo.KlineRepo
}

func NewMarketDomain(mongoClient *database.MongoClient) *MarketDomain {
	return &MarketDomain{
		klineRepo: dao.NewKlineDao(mongoClient.Db),
	}
}

func (m *MarketDomain) SymbolThumbTrend(ctx context.Context, period string, coins []*model.ExchangeCoin) []*model.CoinThumb {
	coinThumbs := make([]*model.CoinThumb, len(coins))
	for i, v := range coins {
		from := tools.ZeroTime()
		end := time.Now().UnixMilli()

		klines, err := m.klineRepo.FindBySymbolTime(ctx, v.Symbol, period, from, end, "")
		if err != nil {
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}
		length := len(klines)
		if length <= 0 {
			coinThumbs[i] = model.DefaultCoinThumb(v.Symbol)
			continue
		}
		trend := make([]float64, len(klines))
		var high float64 = 0
		var low = klines[0].LowestPrice
		var volumes float64 = 0
		var turnover float64 = 0
		for i, v := range klines {
			trend[i] = v.ClosePrice
			if v.HighestPrice > high {
				high = v.HighestPrice
			}
			if v.LowestPrice < low {
				low = v.LowestPrice
			}
			//volumes += v.Volume
			//turnover += v.Turnover
			volumes = op.AddN(volumes, v.Volume, 6)
			turnover = op.AddN(turnover, v.Turnover, 6)
		}
		newKline := klines[0]
		oldKline := klines[len(klines)-1]
		ct := newKline.ToCoinThumb(v.Symbol, oldKline)
		ct.High = high
		ct.Low = low
		ct.Volume = volumes
		ct.Turnover = turnover
		ct.Trend = trend
		coinThumbs[i] = ct
	}
	return coinThumbs
}

func (m *MarketDomain) HistoryKline(ctx context.Context, symbol string, from int64, to int64, period string) ([]*model.History, error) {
	klines, err := m.klineRepo.FindBySymbolTime(ctx, symbol, period, from, to, "asc")
	if err != nil {
		return nil, err
	}
	list := make([]*model.History, len(klines))
	for i, v := range klines {
		h := &model.History{}
		h.Time = v.Time
		h.Open = v.OpenPrice
		h.High = v.HighestPrice
		h.Low = v.LowestPrice
		h.Volume = v.Volume
		h.Close = v.ClosePrice
		list[i] = h
	}
	return list, nil
}

func (d *CoinDomain) FindCoinById(ctx context.Context, id int64) (*model.Coin, error) {
	coin, err := d.CoinRepo.FindById(ctx, id)
	if err != nil {
		logx.Error("查询Coin失败，err=", err)
		return nil, err
	}
	if coin == nil {
		logx.Error("coin为nil，err=", err)
		return nil, errors.New("not support this coin")
	}
	return coin, nil
}
