// Package model
// @Author twilikiss 2024/5/9 22:56:56
package model

type CoinThumb struct {
	Symbol       string    `json:"symbol"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
	Chg          float64   `json:"chg"`
	Change       float64   `json:"change"`
	Volume       float64   `json:"volume"`
	Turnover     float64   `json:"turnover"`
	LastDayClose float64   `json:"lastDayClose"`
	UsdRate      float64   `json:"usdRate"`
	BaseUsdRate  float64   `json:"baseUsdRate"`
	Zone         float64   `json:"zone"`
	DateTime     int64     `json:"dateTime"`
	Trend        []float64 `json:"trend"`
}
type CoinThumbForWs struct {
	Symbol       string    `json:"symbol"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
	Chg          float64   `json:"chg"`
	Change       float64   `json:"change"`
	Volume       float64   `json:"volume"`
	Turnover     float64   `json:"turnover"`
	LastDayClose float64   `json:"lastDayClose"`
	UsdRate      float64   `json:"usdRate"`
	BaseUsdRate  float64   `json:"baseUsdRate"`
	Zone         float64   `json:"zone"`
	DateTime     int64     `json:"dateTime"`
	Trend        []float64 `json:"trend"`
}
