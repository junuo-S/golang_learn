package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println("aaaaaa")
		// 调用了别的函数
		test()
		fmt.Println("bbbbb")
	}()
	// 特地写一个死循环，目的不让主协程结束
	for {

	}
}

func test() {
	defer fmt.Println("cccccc")
	runtime.Goexit() // 终止所在的协程
	fmt.Println("ddddd")
}
