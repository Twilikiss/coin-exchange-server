package dao

import (
	"common/tools"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"jobcenter/internal/model"
	"strconv"
)

type KlineDao struct {
	db *mongo.Database
}

// SaveBatch 存储定时获取的kline数据
func (k *KlineDao) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	ds := make([]interface{}, len(data))
	for i, v := range data {
		ds[i] = v
	}
	_, err := collection.InsertMany(ctx, ds)
	return err
}

func (k *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"time", bson.D{{"$gte", time}}}})
	if err != nil {
		return err
	}
	logx.Infof("%s %s 删除了%d条数据 \n", symbol, period, deleteResult.DeletedCount)
	logx.Info(tools.ToTimeString(tools.ToInt64(strconv.Itoa(int(time)))))
	return nil
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}
