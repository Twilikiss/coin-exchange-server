// Package model
// @Author twilikiss 2024/5/5 1:05:05
package model

type OkxKlineRes struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type OkxConfig struct {
	ApiKey       string
	SecretKey    string
	Pass         string
	Host         string
	Proxy        string
	GetKlineRest string
}
type OkxExchangeRateResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []ExchangeRate `json:"data"`
}
type ExchangeRate struct {
	UsdCny string `json:"usdCny"`
}
