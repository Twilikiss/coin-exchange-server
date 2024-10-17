// Package schedulerTest
// @Author twilikiss 2024/5/4 22:19:19
package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(func() {
		fmt.Println(time.Now().String())
	})
	s.StartBlocking()
}
