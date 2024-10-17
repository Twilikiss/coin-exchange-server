// Package domain
// @Author twilikiss 2024/5/13 23:15:15
package domain

import (
	"common/dbutils"
	"common/op"
	"common/tools"
	"context"
	"errors"
	"exchange/internal/dao"
	"exchange/internal/model"
	"exchange/internal/repo"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/ucclient"
	"time"
)

type ExchangeOrderDomain struct {
	orderRepo repo.OrderRepo
}

func NewExchangeOrderDomain(db *dbutils.ElysiaDB) *ExchangeOrderDomain {
	return &ExchangeOrderDomain{orderRepo: dao.NewOrderDao(db)}
}

func (e *ExchangeOrderDomain) FindOrderHistory(
	ctx context.Context, symbol string, page int64, size int64, memberId int64) ([]*model.ExchangeOrderVo, int64, error) {
	list, total, err := e.orderRepo.FindOrderHistory(ctx, symbol, page, size, memberId)
	if err != nil {
		logx.Error("FindOrderHistory error, err=", err)
		return nil, 0, err
	}
	voList := make([]*model.ExchangeOrderVo, len(list))
	for i, v := range list {
		voList[i] = v.ToVo()
	}
	return voList, total, nil
}
func (e *ExchangeOrderDomain) FindOrderCurrent(
	ctx context.Context, symbol string, page int64, size int64, memberId int64) ([]*model.ExchangeOrderVo, int64, error) {
	list, total, err := e.orderRepo.FindOrderCurrent(ctx, symbol, page, size, memberId)
	if err != nil {
		logx.Error("FindOrderCurrent error, err=", err)
		return nil, 0, err
	}
	voList := make([]*model.ExchangeOrderVo, len(list))
	for i, v := range list {
		voList[i] = v.ToVo()
	}
	return voList, total, nil
}

func (e *ExchangeOrderDomain) FindCurrentTradingCount(
	ctx context.Context, userId int64, symbol string, direction string) (int64, error) {
	return e.orderRepo.FindCurrentTradingCount(ctx, userId, symbol, model.DirectionMap.Code(direction))
}

// AddOrder 返回冻结的金额 -float64
func (e *ExchangeOrderDomain) AddOrder(
	ctx context.Context,
	conn dbutils.DbConn,
	order *model.ExchangeOrder,
	coin *mclient.ExchangeCoin,
	baseWallet *ucclient.MemberWallet,
	coinWallet *ucclient.MemberWallet) (money float64, err error) {

	// 我们的order需要在mysql中创建，并且通过MQ将对于订单信息传递到用户模块，对用户的钱包进行冻结操作
	// 显然在我们钱包完成冻结之前，我们的交易订单都不能直接处于【可交易】状态
	//order.Status = model.Trading
	order.Status = model.Init
	logx.Info("当前的order的status为", order.Status)
	order.TradedAmount = 0
	order.Time = time.Now().UnixMilli()
	order.OrderId = tools.Unq("E")

	// 交易的时候  coin.Fee 费率 手续费 我们做的时候 先不考虑手续费
	// 买 花USDT 市价 price 0 冻结的直接就是amount  卖 BTC
	if order.Direction == model.BUY {
		if order.Type == model.MarketPrice {
			money = order.Amount
		} else {
			//order.Price*order.Amount 精度损失问题
			money = op.MulFloor(order.Price, order.Amount, 8)
		}
		if baseWallet.Balance < money {
			return 0, errors.New("余额不足")
		}
	} else {
		money = order.Amount
		if coinWallet.Balance < money {
			return 0, errors.New("余额不足")
		}
	}
	err = e.orderRepo.Save(ctx, conn, order)
	return money, err
}

func (e *ExchangeOrderDomain) FindOrderByOrderId(ctx context.Context, orderId string) (*model.ExchangeOrder, error) {
	o, err := e.orderRepo.FindOrderByOrderId(ctx, orderId)
	if err != nil {
		logx.Error("findOrderByOrderId has error:", err)
		return nil, errors.New("查询id失败")
	}
	if o == nil {
		// 无法通过OrderId来查询Order
		return nil, errors.New("OrderId不存在")
	}
	return o, nil
}

func (e *ExchangeOrderDomain) UpdateStatusCancel(ctx context.Context, orderId string) error {
	return e.orderRepo.UpdateStatusCancel(ctx, orderId)
}

func (e *ExchangeOrderDomain) UpdateStatusTrading(ctx context.Context, orderId string) error {
	return e.orderRepo.UpdateStatusTrading(ctx, orderId)
}

func (e *ExchangeOrderDomain) FindOrderListBySymbol(ctx context.Context, symbol string, status int) ([]*model.ExchangeOrder, error) {
	return e.orderRepo.FindOrderListBySymbol(ctx, symbol, status)
}

func (e *ExchangeOrderDomain) UpdateOrderComplete(ctx context.Context, order *model.ExchangeOrder) error {
	return e.orderRepo.UpdateOrderComplete(ctx, order.OrderId, order.TradedAmount, order.Turnover, order.Status)
}
