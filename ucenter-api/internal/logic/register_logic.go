package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	logx.Info("api register is working!")
	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()

	reqReq := &register.RegReq{}
	if err := copier.Copy(reqReq, req); err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UCRegisterRPC.RegisterByPhone(ctx, reqReq)
	if err != nil {
		return nil, err
	}
	return
}

func (l *RegisterLogic) SendCode(req *types.CodeRequest) (resp *types.Response, err error) {
	logx.Info("api register is working!")

	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()
	_, err = l.svcCtx.UCRegisterRPC.SendCode(ctx, &register.CodeReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return
}
