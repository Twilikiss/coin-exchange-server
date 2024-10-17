// Package domain
// @Author twilikiss 2024/5/12 18:02:02
package domain

import (
	"common/dbutils"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
)

type CoinDomain struct {
	CoinRepo repo.CoinRepo
}

func NewCoinDomain(db *dbutils.ElysiaDB) *CoinDomain {
	return &CoinDomain{
		CoinRepo: dao.NewCoinDao(db),
	}
}

func (d *CoinDomain) FindCoinInfo(ctx context.Context, unit string) (*model.Coin, error) {
	coin, err := d.CoinRepo.FindByUnit(ctx, unit)
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

func (d *CoinDomain) FindAll(ctx context.Context) (list []*model.Coin, err error) {
	return d.CoinRepo.FindAll(ctx)
}
