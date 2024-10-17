// Package tools
// @Author twilikiss 2024/4/29 1:10:10
package tools

import "k8s.io/apimachinery/pkg/util/rand"

func Gen4Num() int {
	code := rand.IntnRange(1000, 9999)
	return code
}
