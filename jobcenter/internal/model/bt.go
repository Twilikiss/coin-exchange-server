// Package model
// @Author twilikiss 2024/8/10 0:51:51
package model

type BitCoinTransaction struct {
	TxId      string  `bson:"txId"`
	Time      int64   `bson:"time"`
	Value     float64 `bson:"value"`
	BlockHash string  `bson:"blockhash"`
	Address   string  `bson:"address"`
	Type      string  `bson:"type"` // RECHARGE 充值 WITHDRAW 提现
}

const RECHARGE = "RECHARGE"
const WITHDRAW = "WITHDRAW"

func (*BitCoinTransaction) Table() string {
	return "bitcoin_transaction"
}

type BitCoinConfig struct {
	Address string
}
