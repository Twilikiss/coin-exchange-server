// Package logic
// @Author twilikiss 2024/5/16 0:14:14
package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/member"
	"ucenter/internal/domain"
	"ucenter/internal/svc"
)

type MemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain *domain.MemberDomain
	//memberTransactionDomain *domain.MemberTransactionDomain
}

func NewMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLogic {
	return &MemberLogic{
		ctx:          ctx,
		svcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
		memberDomain: domain.NewMemberDomain(svcCtx.Db),
	}
}

func (l *MemberLogic) FindMemberById(req *member.MemberReq) (*member.MemberInfo, error) {
	mem, err := l.memberDomain.FindMemberById(l.ctx, req.MemberId)
	if err != nil {
		logx.Error("FindMemberById error, err=", err)
		return nil, err
	}
	resp := &member.MemberInfo{}
	err = copier.Copy(resp, mem)
	if err != nil {
		logx.Error("转换错误，err=", err)
		return nil, err
	}
	return resp, nil
}
