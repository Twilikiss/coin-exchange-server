// Package logic
// @Author twilikiss 2024/5/3 0:16:16
package logic

import (
	"common/pages"
	"context"
	"errors"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"time"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) History(req *types.ExchangeReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	userId := l.ctx.Value("userId").(int64)
	symbol := req.Symbol
	orderRes, err := l.svcCtx.OrderRPC.FindOrderHistory(ctx, &order.OrderReq{
		Ip:       req.Ip,
		Symbol:   symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		logx.Error("获取历史委托数据失败，err=", err)
		return nil, err
	}
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

func (l *OrderLogic) Current(req *types.ExchangeReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	userId := l.ctx.Value("userId").(int64)
	symbol := req.Symbol
	orderRes, err := l.svcCtx.OrderRPC.FindOrderCurrent(ctx, &order.OrderReq{
		Symbol:   symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

func (l *OrderLogic) AddOrder(req *types.ExchangeReq) (string, error) {
	// 我们需要返回对应的orderId
	// 调用exchange rpc 完成addOrder功能
	value := l.ctx.Value("userId").(int64)
	// 在开始我们的操作前，先简单校验我们的参数是否合法
	if !req.OrderValid() {
		return "", errors.New("参数传递错误")
	}
	orderRes, err := l.svcCtx.OrderRPC.Add(l.ctx, &order.OrderReq{
		Symbol:    req.Symbol,
		UserId:    value,
		Direction: req.Direction,
		Type:      req.Type,
		Price:     req.Price,
		Amount:    req.Amount,
	})
	if err != nil {
		return "", err
	}
	return orderRes.OrderId, nil
}
