package tran

import (
	"common/dbutils"
	"common/dbutils/gorms"
	"gorm.io/gorm"
)

type TransactionImpl struct {
	conn dbutils.DbConn
}

func (t *TransactionImpl) Action(f func(conn dbutils.DbConn) error) error {
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction(db *gorm.DB) *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.New(db),
	}
}
