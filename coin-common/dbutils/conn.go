package dbutils

import "gorm.io/gorm"

type DbConn interface {
	Begin()
	Rollback()
	Commit()
}

type ElysiaDB struct {
	Conn *gorm.DB
}
