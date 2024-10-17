// Package tools
// @Author twilikiss 2024/5/4 18:01:01
package tools

import (
	"fmt"
	"testing"
	"time"
)

func TestGetJwtToken(t *testing.T) {
	token, err := GetJwtToken("elysia", time.Now().Unix(), 4600, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token, err := ParseToken(
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ4MjI0NjksImlhdCI6MTcxNDgxNzg2OSwidXNlcklkIjoxfQ.2YU5MWGZJqzIbrW5z9Zr6CkJmd2N3GmL1Tb7iDir7ZAkyf3VPtk5v8FSEdkh9o3_iOuSYh2_8ZfW0-mOIkBRiQ\n",
		"elysia",
	)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}
