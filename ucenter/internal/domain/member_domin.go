// Package domain
// @Author twilikiss 2024/5/1 14:21:21
package domain

import (
	"common/dbutils"
	"common/tools"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberDomain struct {
	memberRepo repo.MemberRepo
}

func NewMemberDomain(db *dbutils.ElysiaDB) *MemberDomain {
	return &MemberDomain{memberRepo: dao.NewMemberDao(db)}
}

func (m *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	// 设计数据库查询操作
	mem, err := m.memberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return mem, nil
}

func (m *MemberDomain) FindByUserName(ctx context.Context, username string) (*model.Member, error) {
	// 设计数据库查询操作
	mem, err := m.memberRepo.FindByUserName(ctx, username)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return mem, nil
}

func (m *MemberDomain) Register(
	ctx context.Context,
	username string,
	phone string,
	password string,
	country string,
	partner string,
	promotion string) error {
	member := model.NewMember()
	// 对password进行md5加密，注意加盐，不加盐的md5容易被彩虹表进行破解
	// member的字段比较多，但是数据库中这些字段又不能为空，需要为字段加上默认值
	tools.Default(member)
	salt, pwd := tools.Encode(password, nil)

	member.MobilePhone = phone
	member.Username = username
	member.Password = pwd
	member.Country = country
	member.FillSuperPartner(partner)
	member.PromotionCode = promotion
	member.MemberLevel = model.GENERAL
	member.Salt = salt
	member.Avatar = "https://p26-passport.byteacctimg.com/img/user-avatar/264fff3df4d95a25bdcdd2017cf0df5d~80x80.awebp"
	err := m.memberRepo.Save(ctx, member)
	if err != nil {
		logx.Error("数据库异常，error =", err)
		return err
	}
	return nil
}

func (m *MemberDomain) UpdateLoginCountById(ctx context.Context, id int64, step int) error {
	err := m.memberRepo.UpdateLoginCountById(ctx, id, step)
	return err
}

func (m *MemberDomain) FindMemberById(ctx context.Context, memberId int64) (*model.Member, error) {
	id, err := m.memberRepo.FindMemberById(ctx, memberId)
	if err == nil && id == nil {
		return nil, errors.New("用户不存在")
	}
	return id, err
}
