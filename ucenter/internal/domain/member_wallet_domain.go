// Package domain
// @Author twilikiss 2024/5/13 23:15:15
package domain

import (
	"common/dbutils"
	"common/dbutils/tran"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
}

func NewMemberWalletDomain(db *dbutils.ElysiaDB) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
	}
}

func (d *MemberWalletDomain) FindWalletBySymbol(ctx context.Context, id int64, name string, coin *mclient.Coin) (*model.MemberWalletCoin, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, id, name)
	if err != nil {
		return nil, err
	}
	if mw == nil {
		// 如果没找到指定coin的钱包信息，可以直接新建一个该coin类型的钱包并存储
		mw, walletCoin := model.NewMemberWallet(id, coin)
		err := d.memberWalletRepo.Save(ctx, mw)
		if err != nil {
			return nil, err
		}
		return walletCoin, nil
	}
	nwc := &model.MemberWalletCoin{}
	copier.Copy(nwc, mw)
	nwc.Coin = coin
	return nwc, nil
}

func (d *MemberWalletDomain) Freeze(ctx context.Context, conn dbutils.DbConn, userId int64, money float64, symbol string) error {
	// 查询用户钱包
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, userId, symbol)
	if err != nil {
		return err
	}
	if mw.Balance < money {
		// 如果余额不足
		return errors.New("余额不足")
	}
	err = d.memberWalletRepo.UpdateFreeze(ctx, conn, userId, symbol, money)
	if err != nil {
		logx.Error("数据库异常， err=", err)
		return err
	}
	return nil
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoin(ctx context.Context, memberId int64, coinName string) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, memberId, coinName)
	if err != nil {
		logx.Error("数据库异常，err=", err)
		return nil, err
	}
	return mw, nil
}

func (d *MemberWalletDomain) UpdateWalletCoinAndBase(ctx context.Context, baseWallet *model.MemberWallet, coinWallet *model.MemberWallet) error {
	return d.transaction.Action(func(conn dbutils.DbConn) error {
		err := d.memberWalletRepo.UpdateWallet(ctx, conn, baseWallet.Id, baseWallet.Balance, baseWallet.FrozenBalance)
		if err != nil {
			return err
		}
		err = d.memberWalletRepo.UpdateWallet(ctx, conn, coinWallet.Id, coinWallet.Balance, coinWallet.FrozenBalance)
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *MemberWalletDomain) FindWallet(ctx context.Context, id int64) (list []*model.MemberWallet, err error) {
	memberWallets, err := d.memberWalletRepo.FindByMemberId(ctx, id)
	if err != nil {
		return nil, err
	}
	return memberWallets, nil
}

func (d *MemberWalletDomain) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	return d.transaction.Action(func(conn dbutils.DbConn) error {
		err := d.memberWalletRepo.UpdateAddress(ctx, conn, wallet)
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *MemberWalletDomain) GetAllAddress(ctx context.Context, coinName string) ([]string, error) {
	return d.memberWalletRepo.FindAllAddress(ctx, coinName)
}

func (d *MemberWalletDomain) FindByAddress(ctx context.Context, address string) (*model.MemberWallet, error) {
	return d.memberWalletRepo.FindByAddress(ctx, address)
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoinId(ctx context.Context, memberId int64, coinId int64) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinId(ctx, memberId, coinId)
	if err != nil {
		logx.Error("数据库异常，err=", err)
		return nil, err
	}
	return mw, nil
}
