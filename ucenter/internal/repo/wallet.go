// Package repo
// @Author twilikiss 2024/5/13 23:24:24
package repo

import (
	"common/dbutils"
	"context"
	"ucenter/internal/model"
)

type MemberWalletRepo interface {
	Save(ctx context.Context, mw *model.MemberWallet) error
	FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error)
	UpdateFreeze(ctx context.Context, conn dbutils.DbConn, id int64, symbol string, money float64) error
	UpdateWallet(ctx context.Context, conn dbutils.DbConn, id int64, balance float64, frozenBalance float64) error
	FindByMemberId(ctx context.Context, id interface{}) (mvs []*model.MemberWallet, err error)
	UpdateAddress(ctx context.Context, conn dbutils.DbConn, wallet *model.MemberWallet) error
	FindAllAddress(ctx context.Context, coinName string) ([]string, error)
	FindByAddress(ctx context.Context, address string) (*model.MemberWallet, error)
	FindByIdAndCoinId(ctx context.Context, memberId int64, coinId int64) (mw *model.MemberWallet, err error)
}
