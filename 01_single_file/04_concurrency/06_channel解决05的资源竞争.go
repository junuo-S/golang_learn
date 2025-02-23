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

var ch = make(chan int)

func person1() {
	Printer("hello")
	ch <- 666
}

func person2() {
	<-ch // 取数据，没有数据就会阻塞
	Printer("world")
}

func main() {
	go person1()
	go person2()
	// 特地不让主协程结束
	for {

	}
}
