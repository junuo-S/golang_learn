package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(1) // 设置使用的CPU核数，返回值是之前的核数
	fmt.Println("n = ", n)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
