// Package dao
// @Author twilikiss 2024/5/1 22:22:22
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"errors"
	"gorm.io/gorm"
	"market/internal/model"
)

type ExchangeCoinDao struct {
	conn *gorms.GormConn
}

func NewExchangeCoinDao(db *dbutils.ElysiaDB) *ExchangeCoinDao {
	return &ExchangeCoinDao{conn: gorms.New(db.Conn)}
}

func (e *ExchangeCoinDao) FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeCoin{}).Where("visible=?", 1).Find(&list).Error
	return
}

func (e *ExchangeCoinDao) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	session := e.conn.Session(ctx)
	data := &model.ExchangeCoin{}
	err := session.Model(&model.ExchangeCoin{}).Where("symbol=?", symbol).Take(data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有找到
		return nil, nil
	}
	return data, err
}
