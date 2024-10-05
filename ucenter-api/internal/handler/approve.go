// Package handler
// @Author twilikiss 2024/8/12 0:04:04
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

func GetSecuritySetting(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApproveReq
		l := logic.NewApproveLogic(r.Context(), svcCtx)
		resp, err := l.FindSecuritySetting(&req)

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
}
