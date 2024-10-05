// Package model
// @Author twilikiss 2024/5/5 20:21:21
package model

import (
	"common/op"
	"common/tools"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type Kline struct {
	Period       string  `bson:"period,omitempty" json:"period"`
	OpenPrice    float64 `bson:"openPrice,omitempty" json:"openPrice"`
	HighestPrice float64 `bson:"highestPrice,omitempty" json:"highestPrice"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty" json:"lowestPrice"`
	ClosePrice   float64 `bson:"closePrice,omitempty" json:"closePrice"`
	Time         int64   `bson:"time,omitempty" json:"time"`
	Count        float64 `bson:"count,omitempty" json:"count"`       //成交笔数
	Volume       float64 `bson:"volume,omitempty" json:"volume"`     //成交量
	Turnover     float64 `bson:"turnover,omitempty" json:"turnover"` //成交额
	//TimeStr      string  `bson:"timeStr,omitempty"`
}

func NewKline(data []string, period string) *Kline {
	toInt64 := tools.ToInt64(data[0])
	return &Kline{
		Time:         toInt64,
		Period:       period,
		OpenPrice:    tools.ToFloat64(data[1]),
		HighestPrice: tools.ToFloat64(data[2]),
		LowestPrice:  tools.ToFloat64(data[3]),
		ClosePrice:   tools.ToFloat64(data[4]),
		Count:        tools.ToFloat64(data[5]),
		Volume:       tools.ToFloat64(data[6]),
		Turnover:     tools.ToFloat64(data[7]),
		//TimeStr:      tools.ToTimeString(toInt64),
	}
}

// Table BTC/USDT  ETH/USDT  分表
func (*Kline) Table(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

// ToCoinThumb 将Kline数据转换为数据变化趋势
func (k *Kline) ToCoinThumb(symbol string, ct *CoinThumb) *CoinThumb {
	isSame := false
	if ct.Symbol == symbol && ct.DateTime == k.Time {
		// 当我们的DataTime是一样，我们认为这是同一个数据
		isSame = true
	}

	if !isSame {
		newCt := &CoinThumb{}
		err := copier.Copy(newCt, ct)
		if err != nil {
			logx.Error("copy操作失败，err=", err)
			return nil
		}
		newCt.Open = k.OpenPrice
		if newCt.High < k.HighestPrice {
			newCt.High = k.HighestPrice
		}
		//newCt.Low = k.LowestPrice
		if newCt.Low > k.LowestPrice {
			newCt.Low = k.LowestPrice
		}
		newCt.Zone = 0
		newCt.Volume = op.AddN(k.Volume, newCt.Volume, 6)
		newCt.Turnover = op.AddN(k.Turnover, newCt.Turnover, 6)
		newCt.Change = k.LowestPrice - newCt.Close
		newCt.Chg = op.DivN(newCt.Change, newCt.Close, 5)
		newCt.Close = k.ClosePrice
		newCt.UsdRate = k.ClosePrice
		newCt.BaseUsdRate = 1
		newCt.Trend = append(newCt.Trend, k.ClosePrice)
		newCt.DateTime = k.Time
		return newCt
	}
	return ct
}

// InitCoinThumb 初始化
func (k *Kline) InitCoinThumb(symbol string) *CoinThumb {
	ct := &CoinThumb{}
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.High = k.HighestPrice
	ct.Volume = k.Volume
	ct.Turnover = k.Turnover
	ct.Low = k.LowestPrice
	ct.Zone = 0
	ct.Change = 0
	ct.Chg = 0.0
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.Trend = make([]float64, 0)
	ct.DateTime = k.Time
	return ct
}
