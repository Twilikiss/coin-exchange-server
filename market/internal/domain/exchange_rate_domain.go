// Package domain
// @Author twilikiss 2024/5/6 23:59:59
package domain

import "strings"

type ExchangeRateDomain struct {
}

func NewExchangeRateDomain() *ExchangeRateDomain {
	return &ExchangeRateDomain{}
}

func (d *ExchangeRateDomain) UsdRate(unit string) float64 {
	//应该据redis查询，在定时任务做一个根据实际的汇率接口 定期存入redis
	unit = strings.ToUpper(unit)
	if "CNY" == unit {
		return 7.21
	} else if "JPY" == unit {
		return 154
	}
	return 0
}
