// Package bc
// @Author twilikiss 2024/8/3 0:46:46
package bc

import (
	"fmt"
	"testing"
)

func TestBase58Encode(t *testing.T) {
	encode := Base58Encode([]byte("elysia is love"))
	fmt.Println(string(encode))

	data := Base58Decode(encode)
	fmt.Println(string(data))
}
