// Package dao
// @Author twilikiss 2024/8/13 23:52:52
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"ucenter/internal/model"
)

type MemberAddressDao struct {
	conn *gorms.GormConn
}

func (m *MemberAddressDao) FindByMemIdAndCoinId(ctx context.Context, memId int64, coinId int64) (list []*model.MemberAddress, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberAddress{}).
		Where("member_id=? and coin_id=?", memId, coinId).
		Find(&list).Error
	return
}

func NewMemberAddressDao(db *dbutils.ElysiaDB) *MemberAddressDao {
	return &MemberAddressDao{
		conn: gorms.New(db.Conn),
	}
}
