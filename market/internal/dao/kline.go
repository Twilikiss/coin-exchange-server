// Package dao
// @Author twilikiss 2024/5/7 20:24:24
package dao

import (
	"common/tools"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/internal/model"
	"strconv"
)

type KlineDao struct {
	db *mongo.Database
}

// FindBySymbol 按照时间 降序排列
func (k *KlineDao) FindBySymbol(ctx context.Context, symbol, period string, count int64) ([]*model.Kline, error) {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	cur, err := collection.Find(ctx, bson.D{{}}, &options.FindOptions{
		Limit: &count,
		Sort:  bson.D{{"time", -1}}, // 设置按照时间降序排列
	})
	if err != nil {
		return nil, err
	}
	list := make([]*model.Kline, 0)
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// FindBySymbolTime 按照时间范围
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) (list []*model.Kline, err error) {
	mk := &model.Kline{}
	sortInt := -1
	if "asc" == sort {
		sortInt = 1
	}
	collection := k.db.Collection(mk.Table(symbol, period))
	// 这里设置了时间范围并按照时间做了排序处理
	cur, err := collection.Find(ctx,
		bson.D{{"time", bson.D{{"$gte", from}, {"$lte", end}}}},
		&options.FindOptions{
			Sort: bson.D{{"time", sortInt}},
		})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
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
