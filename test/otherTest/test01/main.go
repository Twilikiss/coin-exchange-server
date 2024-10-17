// Package test01
// @Author twilikiss 2024/5/8 1:02:02
package main

import "fmt"

type Student struct {
	Id   string
	Name string
}

func main() {
	list := make([]Student, 0)
	list = append(list, Student{
		Id:   "114514",
		Name: "cxb",
	})
	fmt.Println(list)
}
