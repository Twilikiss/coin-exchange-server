// Package model
// @Author twilikiss 2024/5/13 22:36:36
package model

import (
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
)

type MemberWallet struct {
	Id                int64   `gorm:"column:id"`
	Address           string  `gorm:"column:address"`
	Balance           float64 `gorm:"column:balance"`
	FrozenBalance     float64 `gorm:"column:frozen_balance"`
	ReleaseBalance    float64 `gorm:"column:release_balance"`
	IsLock            int     `gorm:"column:is_lock"`
	MemberId          int64   `gorm:"column:member_id"`
	Version           int     `gorm:"column:version"`
	CoinId            int64   `gorm:"column:coin_id"`
	ToReleased        float64 `gorm:"column:to_released"`
	CoinName          string  `gorm:"column:coin_name"`
	AddressPrivateKey string  `gorm:"address_private_key"`
}

func (*MemberWallet) TableName() string {
	return "member_wallet"
}

func (w *MemberWallet) Copy(coinInfo *mclient.Coin) *MemberWalletCoin {
	mc := &MemberWalletCoin{}
	err := copier.Copy(mc, w)
	if err != nil {
		logx.Error("copy出现错误，err=", err)
	}
	coin := &market.Coin{}
	err = copier.Copy(coin, coinInfo)
	if err != nil {
		logx.Error("copy出现错误，err=", err)
	}
	mc.Coin = coin
	return mc
}

type MemberWalletCoin struct {
	Id             int64        `json:"id" from:"id"`
	Address        string       `json:"address" from:"address"`
	Balance        float64      `json:"balance" from:"balance"`
	FrozenBalance  float64      `json:"frozenBalance" from:"frozenBalance"`
	ReleaseBalance float64      `json:"releaseBalance" from:"releaseBalance"`
	IsLock         int          `json:"isLock" from:"isLock"`
	MemberId       int64        `json:"memberId" from:"memberId"`
	Version        int          `json:"version" from:"version"`
	Coin           *market.Coin `json:"coin" from:"coinId"`
	ToReleased     float64      `json:"toReleased" from:"toReleased"`
}

func NewMemberWallet(memId int64, coin *market.Coin) (*MemberWallet, *MemberWalletCoin) {
	mw := &MemberWallet{
		MemberId: memId,
		CoinId:   int64(coin.Id),
		CoinName: coin.Unit,
	}
	mwc := &MemberWalletCoin{
		MemberId: memId,
		Coin:     coin,
	}
	return mw, mwc
}
