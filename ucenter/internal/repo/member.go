// Package repo
// @Author twilikiss 2024/5/1 14:33:33
package repo

import (
	"context"
	"ucenter/internal/model"
)

type MemberRepo interface {
	// FindByPhone 通过手机号查询用户信息
	FindByPhone(ctx context.Context, phone string) (*model.Member, error)
	Save(ctx context.Context, member *model.Member) error
	FindByUserName(ctx context.Context, username string) (*model.Member, error)
	UpdateLoginCountById(ctx context.Context, id int64, step int) error
	FindMemberById(ctx context.Context, memberId int64) (*model.Member, error)
}
