// Package test01
// @Author twilikiss 2024/6/22 20:27:27
package main

import (
	"fmt"
	"time"
)

// 这个代码可以看出我们test01中使用了goroutine，即便我们的test01已经执行完毕并返回，但是调用test01的main没有退出，所以我们的goroutine运行的test02依旧正常运行
func main() {
	fmt.Println(test01())
	time.Sleep(5 * time.Second)
}

func test01() int {
	fmt.Println("do something!")
	go test02()
	return 100
}

func test02() {
	for {
		fmt.Println("do something1!")
	}
}
