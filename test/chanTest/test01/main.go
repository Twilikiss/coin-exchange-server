// Package test01
// @Author twilikiss 2024/5/8 18:03:03
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Test struct {
	writeChan chan string
}

func NewTest() *Test {
	return &Test{writeChan: make(chan string)}
}

func (c *Test) domain() {
	for {
		select {
		case data := <-c.writeChan:
			fmt.Println(data)
		}
	}
}

func (c *Test) StartWrite() {
	go c.domain()
}

func main() {
	t := NewTest()
	t.StartWrite()

	time.Sleep(10 * time.Second)

	for i := 0; i < 100; i++ {
		t.writeChan <- strconv.Itoa(i)
	}
}
