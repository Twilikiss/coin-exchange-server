// Package handler
// @Author twilikiss 2024/4/28 11:41:41
package handler

import (
	"common"
	"common/tools"
	"errors"
	"google.golang.org/grpc/status"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

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

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
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

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)

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

func SendCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeRequest
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
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
