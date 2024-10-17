// Package domain
// @Author twilikiss 2024/5/5 22:33:33
package domain

import (
	"context"
	"jobcenter/dao"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"jobcenter/internal/repo"
	"log"
)

type KlineDomain struct {
	klineRepo repo.KlineRepo
}

func (d *KlineDomain) SaveBatch(data [][]string, symbol string, period string) {
	klines := make([]*model.Kline, len(data))
	for i, v := range data {
		klines[i] = model.NewKline(v, period) // 将不同组的数组数据构建成对应的kline
	}
	err := d.klineRepo.DeleteGtTime(context.Background(), klines[len(data)-1].Time, symbol, period)
	if err != nil {
		log.Println(err)
		return
	}
	err = d.klineRepo.SaveBatch(context.Background(), klines, symbol, period)
	if err != nil {
		log.Println(err)
	}
}

func NewKlineDomain(cli *database.MongoClient) *KlineDomain {
	return &KlineDomain{
		klineRepo: dao.NewKlineDao(cli.Db),
	}
}
