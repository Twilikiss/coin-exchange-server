// Package tools
// @Author twilikiss 2024/4/30 10:32:32
package tools

import (
	"github.com/thinkeridea/go-extend/exnet"
	"net/http"
)

func GetClientIp(r *http.Request) string {
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	return ip
}
