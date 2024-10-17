// Package handler
// @Author twilikiss 2024/5/2 22:14:14
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

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 简单检查是否包含人机验证信息
		if req.Captcha == nil {
			result := common.NewResult().Deal(nil, errors.New("未包含人机验证参数"), 1001)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		// 在我们进行下一步操作前，先获取一下ip
		req.Ip = tools.GetClientIp(r)

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)

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

func CheckLogin(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.CheckLogin(token)

		if err != nil {
			result := common.NewResult().Deal(resp, err, -1)
			httpx.OkJsonCtx(r.Context(), w, result)
			return
		}

		result := common.NewResult().Deal(resp, nil, -1)
		httpx.OkJsonCtx(r.Context(), w, result)
	}
}
