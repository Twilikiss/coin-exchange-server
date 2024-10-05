// Package dao
// @Author twilikiss 2024/5/12 22:00:00
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"errors"
	"gorm.io/gorm"
	"market/internal/model"
)

type CoinDao struct {
	conn *gorms.GormConn
}

func (d *CoinDao) FindById(ctx context.Context, id int64) (coin *model.Coin, err error) {
	session := d.conn.Session(ctx)
	coin = &model.Coin{}
	err = session.Model(&model.Coin{}).Where("id=?", id).Take(coin).Error
	// 没找到就返回为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return coin, err
}

func (d *CoinDao) FindAll(ctx context.Context) (list []*model.Coin, err error) {
	session := d.conn.Session(ctx)
	err = session.Model(&model.Coin{}).Find(&list).Error
	return
}

func NewCoinDao(db *dbutils.ElysiaDB) *CoinDao {
	return &CoinDao{conn: gorms.New(db.Conn)}
}

func (d *CoinDao) FindByUnit(ctx context.Context, unit string) (*model.Coin, error) {
	session := d.conn.Session(ctx)
	coin := &model.Coin{}
	err := session.Model(&model.Coin{}).Where("unit=?", unit).Take(coin).Error
	// 没找到就返回为空
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return coin, err
}
