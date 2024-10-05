// Package domain
// @Author twilikiss 2024/5/1 14:21:21
package domain

import (
	"common/dbutils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
)

type ExchangeCoinDomain struct {
	exchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeCoinDomain(db *dbutils.ElysiaDB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		exchangeCoinRepo: dao.NewExchangeCoinDao(db),
	}
}

func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) []*model.ExchangeCoin {
	list, err := d.exchangeCoinRepo.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}

func (d *ExchangeCoinDomain) FindBySymbol(req *market.MarketReq, ctx context.Context) (*model.ExchangeCoin, error) {
	symbol := req.GetSymbol()
	exchangeCoin, err := d.exchangeCoinRepo.FindBySymbol(ctx, symbol)
	if err != nil {
		logx.Error("交易对不存在，err=", err)
		return nil, err
	}
	return exchangeCoin, nil
}
