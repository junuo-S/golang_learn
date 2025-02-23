package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("newTask goroutine")
		time.Sleep(time.Second) // 休眠1s
	}
}

func main() {
	go newTask() // 创建goroutine
	for {
		fmt.Println("main goroutine")
		time.Sleep(time.Second) // 休眠1s
	}
}
