// Package dao
// @Author twilikiss 2024/5/1 22:22:22
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"errors"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) UpdateLoginCountById(ctx context.Context, id int64, step int) error {
	session := m.conn.Session(ctx)
	err := session.Exec("update member set login_count = login_count + ? where id = ?", step, id).Error
	return err
}

func (m *MemberDao) Save(ctx context.Context, member *model.Member) error {
	session := m.conn.Session(ctx)
	err := session.Save(member).Error
	return err
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("mobile_phone=?", phone).
		Limit(1).
		Take(&mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) FindByUserName(ctx context.Context, username string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("username=?", username).
		Limit(1).
		Take(&mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return mem, err
}

func (m *MemberDao) FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("id = ?", memberId).Take(&mem).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func NewMemberDao(db *dbutils.ElysiaDB) *MemberDao {
	return &MemberDao{conn: gorms.New(db.Conn)}
}
