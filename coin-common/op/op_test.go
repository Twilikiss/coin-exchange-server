// Package op
// @Author twilikiss 2024/5/11 0:38:38
package op

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	float01 := FloorFloat(2.0/3.0, 8)
	float02 := RoundFloat(2.0/3.0, 8)
	float03 := 2.0 / 3.0
	fmt.Println(float01)
	fmt.Println(float02)
	fmt.Println(float03)
}
