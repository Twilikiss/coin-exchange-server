// Package repo
// @Author twilikiss 2024/5/13 23:24:24
package repo

import (
	"common/dbutils"
	"context"
	"exchange/internal/model"
)

type OrderRepo interface {
	FindOrderHistory(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindOrderCurrent(ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error)
	FindCurrentTradingCount(ctx context.Context, userId int64, symbol string, code int) (int64, error)
	Save(ctx context.Context, conn dbutils.DbConn, order *model.ExchangeOrder) error
	FindOrderByOrderId(ctx context.Context, orderId string) (*model.ExchangeOrder, error)
	UpdateStatusCancel(ctx context.Context, orderId string) error
	UpdateStatusTrading(ctx context.Context, orderId string) error
	FindOrderListBySymbol(ctx context.Context, symbol string, status int) ([]*model.ExchangeOrder, error)
	UpdateOrderComplete(ctx context.Context, orderId string, tradedAmount float64, turnover float64, status int) error
}
