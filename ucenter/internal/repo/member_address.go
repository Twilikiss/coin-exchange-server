// Package repo
// @Author twilikiss 2024/8/13 18:40:40
package repo

import (
	"context"
	"ucenter/internal/model"
)

type MemberAddressRepo interface {
	FindByMemIdAndCoinId(ctx context.Context, memId int64, coinId int64) ([]*model.MemberAddress, error)
}
