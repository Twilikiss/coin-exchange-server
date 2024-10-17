// Package dao
// @Author twilikiss 2024/5/8 0:05:05
package dao

import (
	"common/tools"
	"context"
	"fmt"
	"market/internal/database"
	"testing"
	"time"
)

func TestKlineDao_FindBySymbolTime(t *testing.T) {
	config := database.MongoConfig{
		Url:      "mongodb://localhost:27017",
		Username: "admin",
		Password: "Cxb1314.",
		DataBase: "elcoin",
	}
	mongo := database.ConnectMongo(config)
	dao := NewKlineDao(mongo.Db)
	from := tools.ZeroTime()
	end := time.Now().UnixMilli()
	klines, err := dao.FindBySymbolTime(context.Background(), "BTC/USDT", "1H", from, end, "")
	if err != nil {
		t.Error(err)
	}
	for _, v := range klines {
		fmt.Println(v.OpenPrice)
	}
}
