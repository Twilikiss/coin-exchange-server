// Package domain
// @Author twilikiss 2024/8/13 23:55:55
package domain

import (
	"common/dbutils"
	"context"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberTAddressDomain struct {
	memberAddressRepo repo.MemberAddressRepo
}

func (d *MemberTAddressDomain) FindAddressList(ctx context.Context, userId int64, coinId int64) ([]*model.MemberAddress, error) {
	return d.memberAddressRepo.FindByMemIdAndCoinId(ctx, userId, coinId)
}

func NewMemberTAddressDomain(db *dbutils.ElysiaDB) *MemberTAddressDomain {
	return &MemberTAddressDomain{
		memberAddressRepo: dao.NewMemberAddressDao(db),
	}
}
