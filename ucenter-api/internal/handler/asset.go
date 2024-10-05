// Package handler
// @Author twilikiss 2024/5/13 22:46:46
package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

func FindWalletBySymbolHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetReq

		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 在我们进行下一步操作前，先获取一下ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewAssetLogic(r.Context(), svcCtx)
		resp, err := l.FindWalletBySymbol(&req)

		// =================================下面是通用处理RPC结果的结构======================================
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
	// ==============================================END===============================================
}

func FindWallet(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetReq

		// 在我们进行下一步操作前，先获取一下ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewAssetLogic(r.Context(), svcCtx)
		resp, err := l.FindWallet(&req)

		// =================================下面是通用处理RPC结果的结构======================================
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
	// ==============================================END===============================================
}

func RestAddress(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetReq

		// 处理我们的表单数据
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 在我们进行下一步操作前，先获取一下ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewAssetLogic(r.Context(), svcCtx)
		resp, err := l.RestAddress(&req)

		// =================================下面是通用处理RPC结果的结构======================================
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
	// ==============================================END===============================================
}

func FindAllTransaction(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AssetReq

		// 处理我们的表单数据
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 在我们进行下一步操作前，先获取一下ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewAssetLogic(r.Context(), svcCtx)
		resp, err := l.FindAllTransaction(&req)

		// =================================下面是通用处理RPC结果的结构======================================
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
	// ==============================================END===============================================
}
