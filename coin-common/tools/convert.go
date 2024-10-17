// Package tools
// @Author twilikiss 2024/5/5 20:26:26
package tools

import (
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

func ToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		logx.Error(err)
		return 0
	}
	return i
}

func ToFloat64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		logx.Error(err)
		return 0
	}
	return i
}
