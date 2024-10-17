// Package consumer
// @Author twilikiss 2024/8/15 0:11:11
package consumer

import (
	"common/dbutils"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
	"ucenter/internal/model"
)

func WithdrawConsumer(kafkaCli *database.KafkaClient, db *dbutils.ElysiaDB, address string) {
	//获取到提现记录
	//创建BTC网络交易
	//要将交易发送到BTC网络 这时候经过矿工的打包之后 全球可见
	//创建交易的时候 一定要有手续费
	//UTXO unspend 地址的余额  -> 交易的input
	withdrawDomain := domain.NewWithdrawDomain(db, nil, address)
	for {
		kafkaData := kafkaCli.Read()
		var wr model.WithdrawRecord
		json.Unmarshal(kafkaData.Data, &wr)
		ctx := context.Background()
		err := withdrawDomain.Withdraw(ctx, wr)
		if err != nil {
			logx.Error("提币交易失败，err=", err)
			kafkaCli.RPut(kafkaData)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
