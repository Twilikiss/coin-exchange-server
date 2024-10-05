// Package handler
// @Author twilikiss 2024/5/7 13:18:18
package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"market-api/internal/logic"
	"market-api/internal/svc"
	"market-api/internal/types"
	"net/http"
)

func NewSymbolThumbTrendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewSymbolThumbTrendHandler")

		var req types.MarketReq

		// 因为我们采用的是动态路由传参
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewMarketLogic(r.Context(), svcCtx)
		resp, err := l.SymbolThumbTrend(&req)

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

func NewSymbolThumbHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewSymbolThumbHandler")

		var req types.MarketReq

		// 因为我们采用的是动态路由传参
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewMarketLogic(r.Context(), svcCtx)
		resp, err := l.SymbolThumb(&req)

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

func NewSymbolInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewSymbolInfoHandler")

		var req types.MarketReq

		// 这里我们接收的表单类型
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewMarketLogic(r.Context(), svcCtx)
		resp, err := l.SymbolInfo(&req)

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

func NewCoinInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewCoinInfoHandler")

		var req types.MarketReq

		// 这里我们接收的表单类型
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewMarketLogic(r.Context(), svcCtx)
		resp, err := l.CoinInfo(&req)

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

func NewHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewCoinInfoHandler")

		var req types.MarketReq

		// 这里我们接收的表单类型
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewMarketLogic(r.Context(), svcCtx)
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
