// Package server
// @Author twilikiss 2024/8/14 0:12:12
package server

import (
	"context"
	"grpc-common/ucenter/types/withdraw"
	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type WithdrawServer struct {
	svcCtx *svc.ServiceContext
	withdraw.UnimplementedWithdrawServer
}

func NewWithdrawServer(svcCtx *svc.ServiceContext) *WithdrawServer {
	return &WithdrawServer{
		svcCtx: svcCtx,
	}
}

func (w *WithdrawServer) FindAddressByCoinId(ctx context.Context, wq *withdraw.WithdrawReq) (*withdraw.AddressSimpleList, error) {
	l := logic.NewWithdrawLogic(ctx, w.svcCtx)
	return l.FindAddressByCoinId(wq)
}
func (w *WithdrawServer) SendCode(ctx context.Context, in *withdraw.WithdrawReq) (*withdraw.NoRes, error) {
	l := logic.NewWithdrawLogic(ctx, w.svcCtx)
	return l.SendCode(in)
}

func (w *WithdrawServer) WithdrawCode(ctx context.Context, in *withdraw.WithdrawReq) (*withdraw.NoRes, error) {
	l := logic.NewWithdrawLogic(ctx, w.svcCtx)
	return l.WithdrawCode(in)
}

func (s *WithdrawServer) WithdrawRecord(ctx context.Context, in *withdraw.WithdrawReq) (*withdraw.RecordList, error) {
	l := logic.NewWithdrawLogic(ctx, s.svcCtx)
	return l.WithdrawRecord(in)
}
