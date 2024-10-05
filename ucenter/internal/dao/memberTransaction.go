// Package dao
// @Author twilikiss 2024/8/4 1:10:10
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"common/tools"
	"context"
	"errors"
	"gorm.io/gorm"
	"ucenter/internal/model"
)

type MemberTransactionDao struct {
	conn *gorms.GormConn
}

func (d *MemberTransactionDao) Save(ctx context.Context, transaction *model.MemberTransaction) error {
	session := d.conn.Session(ctx)
	err := session.Save(transaction).Error
	return err
}

func (d *MemberTransactionDao) FindByAmountAndTime(ctx context.Context, address string, value float64, time int64) (mt *model.MemberTransaction, err error) {
	session := d.conn.Session(ctx)
	err = session.Model(&model.MemberTransaction{}).
		Where("address=? and amount=? and create_time=?", address, value, time).
		Limit(1).Take(&mt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (d *MemberTransactionDao) FindTransaction(ctx context.Context, pageNo int, pageSize int, memberId int64, startTime string, endTime string, symbol string, transactionType string) (list []*model.MemberTransaction, total int64, err error) {
	session := d.conn.Session(ctx)
	db := session.Model(&model.MemberTransaction{}).Where("member_id=?", memberId)
	if transactionType != "" {
		db.Where("type=?", tools.ToInt64(transactionType))
	}
	if startTime != "" && endTime != "" {
		sTime := tools.ToMill(startTime)
		eTime := tools.ToMill(endTime)
		db.Where("create_time >= ? and create_time <= ?", sTime, eTime)
	}
	if symbol != "" {
		db.Where("symbol=?", symbol)
	}
	offset := (pageNo - 1) * pageSize
	db.Count(&total)
	db.Order("create_time desc").Offset(offset).Limit(pageSize)
	err = db.Find(&list).Error
	return
}

func NewMemberTransactionDao(db *dbutils.ElysiaDB) *MemberTransactionDao {
	return &MemberTransactionDao{conn: gorms.New(db.Conn)}
}
