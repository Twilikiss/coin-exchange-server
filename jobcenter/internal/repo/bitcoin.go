// Package repo
// @Author twilikiss 2024/8/10 0:50:50
package repo

import "jobcenter/internal/model"

type BtcTransactionRepo interface {
	FindByTxId(txId string) (*model.BitCoinTransaction, error)
	Save(bt *model.BitCoinTransaction) error
}
