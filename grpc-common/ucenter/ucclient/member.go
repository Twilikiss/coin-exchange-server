// Package ucclient
// @Author twilikiss 2024/5/16 0:08:08
package ucclient

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"grpc-common/ucenter/types/member"
)

type (
	MemberReq  = member.MemberReq
	MemberInfo = member.MemberInfo

	Member interface {
		FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error)
	}

	defaultMember struct {
		cli zrpc.Client
	}
)

func NewMember(cli zrpc.Client) Member {
	return &defaultMember{
		cli: cli,
	}
}

func (m *defaultMember) FindMemberById(ctx context.Context, in *MemberReq, opts ...grpc.CallOption) (*MemberInfo, error) {
	client := member.NewMemberClient(m.cli.Conn())
	return client.FindMemberById(ctx, in, opts...)
}
