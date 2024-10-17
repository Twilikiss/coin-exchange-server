// Package handler
// @Author twilikiss 2024/4/28 11:41:41
package handler

import (
	"common"
	"common/tools"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"market-api/internal/logic"
	"market-api/internal/svc"
	"market-api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

//type RegisterHandler struct {
//	svcCtx *svc.ServiceContext
//}
//
//func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
//	return &RegisterHandler{
//		svcCtx: svcCtx,
//	}
//}

func NewUsdRateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("消息来到了NewUsdRateHandler")
		var req types.RateRequest

		// 因为我们采用的是动态路由传参
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取远程访问的ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewExchangeRateLogic(r.Context(), svcCtx)
		resp, err := l.UsdRate(&req)

		if err != nil {
			// 简单处理一下RPC相关错误
			s, _ := status.FromError(err)

			result := common.NewResult().Deal(nil, errors.New(s.Message()), -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp.Rate, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}
