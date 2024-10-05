// Package dao
// @Author twilikiss 2024/5/13 23:30:30
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"errors"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemberWalletDao struct {
	conn *gorms.GormConn
}

func (m *MemberWalletDao) FindByIdAndCoinId(ctx context.Context, memberId int64, coinId int64) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_id=?", memberId, coinId).
		Take(&mw).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (m *MemberWalletDao) FindByAddress(ctx context.Context, address string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("address=?", address).Take(&mw).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (m *MemberWalletDao) FindAllAddress(ctx context.Context, coinName string) ([]string, error) {
	session := m.conn.Session(ctx)
	var list []string
	err := session.Model(&model.MemberWallet{}).Where("coin_name=?", coinName).Select("address").Find(&list).Error
	return list, err
}

func (m *MemberWalletDao) UpdateAddress(ctx context.Context, conn dbutils.DbConn, wallet *model.MemberWallet) error {
	gormConn := conn.(*gorms.GormConn)
	tx := gormConn.Tx(ctx)
	//Update
	updateSql := "update member_wallet set address=?,address_private_key=? where id=?"
	err := tx.Model(&model.MemberWallet{}).Exec(updateSql, wallet.Address, wallet.AddressPrivateKey, wallet.Id).Error
	return err
}

func (m *MemberWalletDao) FindByMemberId(ctx context.Context, id interface{}) (mvs []*model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).Where("member_id = ?", id).Find(&mvs).Error
	return
}

func (m *MemberWalletDao) UpdateWallet(ctx context.Context, conn dbutils.DbConn, id int64, balance float64, frozenBalance float64) error {
	gormConn := conn.(*gorms.GormConn)
	tx := gormConn.Tx(ctx)
	//Update
	updateSql := "update member_wallet set balance=?,frozen_balance=? where id=?"
	err := tx.Model(&model.MemberWallet{}).Exec(updateSql, balance, frozenBalance, id).Error
	return err
}

func (m *MemberWalletDao) UpdateFreeze(ctx context.Context, conn dbutils.DbConn, memberId int64, symbol string, money float64) error {
	con := conn.(*gorms.GormConn)
	session := con.Tx(ctx)
	sql := "update member_wallet set balance=balance-?, frozen_balance=frozen_balance+? where member_id=? and coin_name=?"
	err := session.Model(&model.MemberWallet{}).Exec(sql, money, money, memberId, symbol).Error
	return err
}

func NewMemberWalletDao(db *dbutils.ElysiaDB) *MemberWalletDao {
	return &MemberWalletDao{
		conn: gorms.New(db.Conn),
	}
}

func (m *MemberWalletDao) Save(ctx context.Context, mw *model.MemberWallet) error {
	session := m.conn.Session(ctx)
	err := session.Save(&mw).Error
	return err
}

func (m *MemberWalletDao) FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.MemberWallet{}).
		Where("member_id=? and coin_name=?", memId, coinName).
		Take(&mw).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}
