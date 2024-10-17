// Package logic
// @Author twilikiss 2024/8/1 18:06:06
package logic

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"jobcenter/internal/model"
	"log"
	"sync"
	"time"
)

type Rate struct {
	wg sync.WaitGroup
	c  model.OkxConfig
	ch *redis.Redis
}

func (r *Rate) Do() {
	//要获取 人民币对美元的汇率
	r.wg.Add(1)
	go r.CnyUsdRate()
	r.wg.Wait()
}

func (r *Rate) CnyUsdRate() {
	//请求接口 获取到最新的汇率 存入redis即可
	//发起http请求 获取数据
	api := r.c.Host + "/api/v5/market/exchange-rate"
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/exchange-rate", r.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = r.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = r.c.Pass
	resp, err := tools.GetWithHeader(
		api,
		header,
		r.c.Proxy,
	)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	var result = &model.OkxExchangeRateResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	cny := result.Data[0].UsdCny
	//存入redis
	err = r.ch.Set("USDT::CNY::RATE", cny)
	if err != nil {
		logx.Error("存入redis失败，err=", err)
	}
	r.wg.Done()
}

func NewRate(c model.OkxConfig, cache *redis.Redis) *Rate {
	return &Rate{
		c:  c,
		ch: cache,
	}
}
