package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, ch := range str {
		fmt.Printf(string(ch))
		time.Sleep(time.Second)
	}
	fmt.Println()
}

func main() {
	go Printer("hello")
	go Printer("world")
	// 特地不让主协程结束
	for {

	}
}
