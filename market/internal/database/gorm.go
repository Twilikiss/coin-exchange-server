package database

import (
	"common/dbutils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlConfig struct {
	DataSource string
}

func ConnMysql(c MysqlConfig) *dbutils.ElysiaDB {
	var err error
	_db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db, _ := _db.DB()
	//连接池配置
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return &dbutils.ElysiaDB{
		Conn: _db,
	}
}
