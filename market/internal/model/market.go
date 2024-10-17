// Package model
// @Author twilikiss 2024/5/7 20:27:27
package model

type CoinThumb struct {
	Symbol       string    `json:"symbol,omitempty"`
	Open         float64   `json:"open,omitempty"`
	High         float64   `json:"high,omitempty"`
	Low          float64   `json:"low,omitempty"`
	Close        float64   `json:"close,omitempty"`
	Chg          float64   `json:"chg,omitempty"`
	Change       float64   `json:"change,omitempty"`
	Volume       float64   `json:"volume,omitempty"`
	Turnover     float64   `json:"turnover,omitempty"`
	LastDayClose float64   `json:"lastDayClose,omitempty"`
	UsdRate      float64   `json:"usdRate,omitempty"`
	BaseUsdRate  float64   `json:"baseUsdRate,omitempty"`
	Zone         float64   `json:"zone,omitempty"`
	DateTime     int64     `json:"dateTime,omitempty"`
	Trend        []float64 `json:"trend,omitempty"`
}

type History struct {
	Time   int64   `json:"time,omitempty"`
	Open   float64 `json:"open,omitempty"`
	Close  float64 `json:"close,omitempty"`
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Volume float64 `json:"volume,omitempty"`
}
