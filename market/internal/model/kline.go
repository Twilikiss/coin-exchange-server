// Package model
// @Author twilikiss 2024/5/7 20:33:33
package model

import "common/op"

type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
}

func (*Kline) Table(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

func (k *Kline) ToCoinThumb(symbol string, end *Kline) *CoinThumb {
	ct := &CoinThumb{}
	ct.Symbol = symbol
	// 针对websocket
	ct.Close = end.ClosePrice
	ct.Open = end.OpenPrice
	ct.Zone = 0
	ct.Change = k.ClosePrice - end.ClosePrice
	// TODO 添加对应的op操作
	ct.Chg = op.DivN(ct.Change, end.ClosePrice, 5)
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.DateTime = k.Time
	return ct
}

func DefaultCoinThumb(symbol string) *CoinThumb {
	ct := &CoinThumb{}
	ct.Symbol = symbol
	ct.Trend = []float64{}
	return ct
}
