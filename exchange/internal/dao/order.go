// Package dao
// @Author twilikiss 2024/5/13 23:30:30
package dao

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"context"
	"errors"
	"exchange/internal/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	conn *gorms.GormConn
}

func (e *OrderDao) UpdateOrderComplete(ctx context.Context, orderId string, tradedAmount float64, turnover float64, status int) error {
	session := e.conn.Session(ctx)
	updateSql := "update exchange_order set traded_amount=?,turnover=?,status=? where order_id=? and status=?"
	err := session.Model(&model.ExchangeOrder{}).Exec(updateSql, tradedAmount, turnover, status, orderId, model.Trading).Error
	return err
}

func (e *OrderDao) FindOrderListBySymbol(ctx context.Context, symbol string, status int) (list []*model.ExchangeOrder, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol=? and status=?", symbol, status).Find(&list).Error
	return
}

func (e *OrderDao) UpdateStatusTrading(ctx context.Context, orderId string) error {
	session := e.conn.Session(ctx)
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id = ?", orderId).Update("status", model.Trading).Error
	return err
}

func (e *OrderDao) UpdateStatusCancel(ctx context.Context, orderId string) error {
	session := e.conn.Session(ctx)
	err := session.Model(&model.ExchangeOrder{}).
		Where("order_id = ?", orderId).Update("status", model.Canceled).Error
	return err
}

func (e *OrderDao) FindOrderByOrderId(ctx context.Context, orderId string) (order *model.ExchangeOrder, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("order_id = ?", orderId).
		Take(&order).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

func (e *OrderDao) Save(ctx context.Context, conn dbutils.DbConn, order *model.ExchangeOrder) error {
	c := conn.(*gorms.GormConn)
	tx := c.Tx(ctx)
	err := tx.Save(&order).Error
	return err
}

func (e *OrderDao) FindCurrentTradingCount(ctx context.Context, userId int64, symbol string, code int) (total int64, err error) {
	session := e.conn.Session(ctx)
	session.Model(&model.ExchangeOrder{}).
		Where("symbol = ? and member_id = ? and direction = ? and status = ?", symbol, userId, code, model.Trading).
		Count(&total)
	return
}

func NewOrderDao(db *dbutils.ElysiaDB) *OrderDao {
	return &OrderDao{
		conn: gorms.New(db.Conn),
	}
}

func (e *OrderDao) FindOrderHistory(
	ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol = ? and member_id = ?", symbol, memberId).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&list).Error
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol = ? and member_id = ?", symbol, memberId).
		Count(&total).Error
	return
}
func (e *OrderDao) FindOrderCurrent(
	ctx context.Context, symbol string, page int64, size int64, memberId int64) (list []*model.ExchangeOrder, total int64, err error) {
	session := e.conn.Session(ctx)
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol = ? and member_id = ? and status = ?", symbol, memberId, model.Trading).
		Limit(int(size)).
		Offset(int((page - 1) * size)).
		Find(&list).Error
	err = session.Model(&model.ExchangeOrder{}).
		Where("symbol = ? and member_id = ?", symbol, memberId).
		Count(&total).Error
	return
}
