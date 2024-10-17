// Package handler
// @Author twilikiss 2024/8/13 0:05:05
package handler

import (
	"common"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

func QueryWithdrawCoin(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawReq

		l := logic.NewWithdrawLogic(r.Context(), svcCtx)
		resp, err := l.QueryWithdrawCoin(&req)
		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(resp, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func SendCode(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawReq

		l := logic.NewWithdrawLogic(r.Context(), svcCtx)
		resp, err := l.SendCode(&req)
		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(resp, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func WithdrawCode(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewWithdrawLogic(r.Context(), svcCtx)
		resp, err := l.WithdrawCode(&req)
		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(resp, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func Record(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewWithdrawLogic(r.Context(), svcCtx)
		resp, err := l.Record(&req)
		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(resp, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}
