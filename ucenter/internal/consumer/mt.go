// Package consumer
// @Author twilikiss 2024/8/10 1:28:28
package consumer

import (
	"common/dbutils"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
)

type BitCoinTransactionResult struct {
	Value   float64 `json:"value"`
	Time    int64   `json:"time"`
	Address string  `json:"address"`
	Type    string  `json:"type"`
	Symbol  string  `json:"symbol"`
}

func BitCoinTransaction(redisCli *redis.Redis, kafkaCli *database.KafkaClient, db *dbutils.ElysiaDB) {
	for {
		kafkaData := kafkaCli.Read()
		var bt BitCoinTransactionResult
		_ = json.Unmarshal(kafkaData.Data, &bt)
		//解析出来数据 调用domain存储到数据库即可
		transactionDomain := domain.NewMemberTransactionDomain(db)
		err := transactionDomain.SaveRecharge(bt.Address, bt.Value, bt.Time, bt.Type, bt.Symbol)

		// TODO 存储完对应交易后，需要调用gorm，将我们的充值的金额加入到我们的钱包中

		if err != nil {
			time.Sleep(200 * time.Millisecond)
			kafkaCli.RPut(kafkaData)
		}
	}
}
