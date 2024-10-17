// Package repo
// @Author twilikiss 2024/5/1 14:33:33
package repo

import (
	"context"
	"market/internal/model"
)

type ExchangeCoinRepo interface {
	FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error)
	FindBySymbol(ctx context.Context, symbol string) (exchangeCoin *model.ExchangeCoin, err error)
}
