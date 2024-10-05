// Package domain
// @Author twilikiss 2024/4/30 0:01:01
package domain

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type vaptchaReq struct {
	Id        string `json:"id"`
	SecretKey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

type CaptchaDomain struct {
}

func (d *CaptchaDomain) Verify(server string, vid string, key string, token string, scene int, ip string) bool {
	// 发送对应的服务端二次验证请求
	post, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		SecretKey: key,
		Scene:     scene,
		Token:     token,
		Ip:        ip,
	})
	if err != nil {
		logx.Errorf("发生致命错误，err=%s\n", err)
		return false
	}
	var result vaptchaRsp
	err = json.Unmarshal(post, &result)
	if err != nil {
		logx.Errorf("发生致命错误，err=%s\n", err)
		return false
	}
	logx.Infof("人机验证结果为：success=%d, score=%d, msg=%s\n", result.Success, result.Score, result.Msg)
	return result.Success == 1
}

func NewCaptchaDomain() *CaptchaDomain {
	return &CaptchaDomain{}
}
