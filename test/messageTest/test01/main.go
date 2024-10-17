// Package test01
// @Author twilikiss 2024/4/29 8:42:42
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type MessageResponse struct {
	Code   int    `json:"error_code"`
	Reason string `json:"reason"`
	Result string `json:"result"`
	Sn     string `json:"sn"`
}

func main() {
	params := url.Values{}
	params.Add("mobile", "13104912312")
	s := "114514"
	params.Add("content", "【微云信息团队】您的验证码为："+s+",有效时长为5分钟，若非本人操作，请忽略.")
	params.Add("key", "f383fb2409144c728c1ce790a0df7962")

	data := "http://apis.haoservice.com/sms/sendv2?" + params.Encode()

	req, _ := http.NewRequest("GET", data, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var messageResponse MessageResponse

	err := json.Unmarshal(body, &messageResponse)
	if err != nil {
		log.Println(err)
	}
	log.Printf("err_code=%d,status=%s,result=%s,sn=%s",
		messageResponse.Code,
		messageResponse.Reason,
		messageResponse.Result,
		messageResponse.Sn,
	)
}
