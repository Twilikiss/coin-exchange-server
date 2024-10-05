// Package repo
// @Author twilikiss 2024/5/7 20:25:25
package repo

import (
	"context"
	"market/internal/model"
)

type KlineRepo interface {
	FindBySymbol(ctx context.Context, symbol, period string, count int64) ([]*model.Kline, error)
	FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) ([]*model.Kline, error)
}
