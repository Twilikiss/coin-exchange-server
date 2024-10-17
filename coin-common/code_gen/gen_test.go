// Package code_gen
// @Author twilikiss 2024/5/11 11:21:21
package code_gen

import "testing"

func TestGenStruct(t *testing.T) {
	GenStruct("coin", "Coin")
}

func TestGenProtoMessage(t *testing.T) {
	GenProtoMessage("member", "Member")
}
