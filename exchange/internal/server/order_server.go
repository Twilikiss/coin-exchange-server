// Package server
// @Author twilikiss 2024/5/13 23:08:08
package server

import (
	"context"
	"exchange/internal/logic"
	"exchange/internal/svc"
	"grpc-common/exchange/types/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (o *OrderServer) FindOrderHistory(ctx context.Context, in *order.OrderReq) (*order.OrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, o.svcCtx)
	return l.FindOrderHistory(in)
}

func (o *OrderServer) FindOrderCurrent(ctx context.Context, in *order.OrderReq) (*order.OrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, o.svcCtx)
	return l.FindOrderCurrent(in)
}

func (o *OrderServer) Add(ctx context.Context, in *order.OrderReq) (*order.AddOrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, o.svcCtx)
	return l.Add(in)
}
func (o *OrderServer) FindByOrderId(ctx context.Context, req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	l := logic.NewExchangeOrderLogic(ctx, o.svcCtx)
	return l.FindByOrderId(req)
}
func (o *OrderServer) CancelOrder(ctx context.Context, req *order.OrderReq) (*order.CancelOrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, o.svcCtx)
	return l.CancelOrder(req)
}
