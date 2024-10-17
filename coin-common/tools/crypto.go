// Package tools
// @Author twilikiss 2024/5/5 0:30:30
package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// ComputeHmacSha256 输入需要被加密的文本和密钥，返回加密后的字串
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
