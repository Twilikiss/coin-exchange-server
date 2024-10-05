// Package tools
// @Author twilikiss 2024/8/14 23:53:53
package tools

import (
	"fmt"
	"testing"
)

func TestToTimeString(t *testing.T) {
	timeString := ToTimeString(1723632282000)
	fmt.Println(timeString)
}
