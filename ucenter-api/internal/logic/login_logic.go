// Package logic
// @Author twilikiss 2024/5/3 0:16:16
package logic

import (
	"common/tools"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/login"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	logx.Info("api register is working!")
	ctx, cancelFunc := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancelFunc()

	loginRep := &login.LoginReq{}
	if err := copier.Copy(loginRep, req); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	loginResp, err := l.svcCtx.UCLoginRPC.Login(ctx, loginRep)

	if err != nil {
		return nil, err
	}

	data := &types.LoginRes{}

	if err := copier.Copy(data, loginResp); err != nil {
		logx.Error("copy时出现错误，error=", err)
		return nil, errors.New("内部错误，请联系管理员")
	}

	return data, nil
}

func (l *LoginLogic) CheckLogin(token string) (isLogin bool, err error) {
	if token == "" {
		return false, nil
	}
	secret := l.svcCtx.Config.JWT.AccessSecret
	_, err = tools.ParseToken(token, secret)
	if err != nil {
		logx.Error("token校验失败，error=", err)
		return false, nil
	}
	return true, nil
}
