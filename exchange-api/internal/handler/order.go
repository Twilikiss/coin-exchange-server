// Package handler
// @Author twilikiss 2024/5/15 0:29:29
package handler

import (
	"common"
	"common/tools"
	"errors"
	"exchange-api/internal/logic"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

func NewOrderHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewOrderHistoryHandler")
		var req types.ExchangeReq

		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.History(&req)

		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(nil, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func NewOrderCurrentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewOrderCurrentHandler")
		var req types.ExchangeReq

		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.Current(&req)

		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(nil, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}

func NewAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewAddHandler")
		var req types.ExchangeReq

		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrder(&req)

		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(nil, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}
