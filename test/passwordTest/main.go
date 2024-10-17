// Package passwordTest
// @Author twilikiss 2024/5/2 20:36:36
package main

import (
	"common/tools"
	"fmt"
)

func main() {
	salt, pwd := tools.Encode("114514", nil)
	fmt.Println(salt)
	fmt.Println("******************")
	fmt.Println(pwd)
}
