// Package tools
// @Author twilikiss 2024/5/17 18:22:22
package tools

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

// TODO Unq 后面OrderId的生成考虑采用雪花算法

// Unq prefix+（毫秒值和随机值相结合）
func Unq(prefix string) string {
	milli := time.Now().UnixMilli()
	intn := rand.IntnRange(100000, 999999)
	return fmt.Sprintf("%s%d%d", prefix, milli, intn)
}
