// Package logic
// @Author twilikiss 2024/5/4 23:52:52
package logic

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"jobcenter/internal/domain"
	"jobcenter/internal/model"
	"jobcenter/internal/svc"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Kline struct {
	wg          sync.WaitGroup
	ctx         *svc.ServiceContext
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
	ch          *redis.Redis
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	// 获取某个币的k线数据 BTC-USDT 和 ETH-USDT
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId string, symbol string, period string) {
	api := "GET/api/v5/market/candles?instId=" + instId + "&bar=" + period
	// 发起http请求获取对应数据
	timestamp := tools.ISO(time.Now())

	secretKey := k.ctx.Config.Okx.SecretKey
	sign := tools.ComputeHmacSha256(timestamp+api, secretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.ctx.Config.Okx.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.ctx.Config.Okx.Pass

	path := k.ctx.Config.Okx.Host + k.ctx.Config.Okx.GetKlineRest
	u := url.Values{}
	u.Add("instId", instId)
	u.Add("bar", period)
	path = path + "?" + u.Encode()

	respBody, err := tools.GetWithHeader(
		path,
		header,
		k.ctx.Config.Okx.Proxy,
	)
	if err != nil {
		logx.Info(err)
	} else {
		//log.Println(instId, string(respBody))
		resp := &model.OkxKlineRes{}
		err := json.Unmarshal(respBody, resp)
		if err != nil {
			logx.Error(err)
		} else {
			if resp.Code == "0" {
				logx.Info("++++++++++++++获取到K线数据++++++++++++")
				//代表成功
				k.klineDomain.SaveBatch(resp.Data, symbol, period)
				//logx.Info("itsId:", instId, "period:", period)
				//logx.Info("result kline result:", resp)

				if period == "1m" {
					//把这个最新的数据result.Data[0] 推送到market服务，推送到前端页面，实时进行变化
					//->kafka->market kafka消费者进行数据消费-> 通过websocket通道发送给前端 ->前端更新数据
					data := resp.Data[0]
					k.queueDomain.Send1mKline(data, symbol)

					//放入redis 将其最新的价格
					key := strings.ReplaceAll(instId, "-", "::")
					err = k.ch.Set(key+"::RATE", data[4]) // 把通过api获取到的指定汇率存入redis
					if err != nil {
						logx.Error("存入redis失败，err=", err)
					}
				}
				logx.Info("+++++++++++++END++++++++++++")
			}
		}
	}
	k.wg.Done()
}

func NewKline(svcCtx *svc.ServiceContext) *Kline {
	return &Kline{
		ctx:         svcCtx,
		klineDomain: domain.NewKlineDomain(svcCtx.MongoClient),
		queueDomain: domain.NewQueueDomain(svcCtx.KafkaClient),
		ch:          svcCtx.Cache,
	}
}
