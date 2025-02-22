package main

import (
	"fmt"
	"os"
)

var a = 100

func main() {
	// 获取命令行参数
	// os.Args[0] 是命令行的名称
	// os.Args[1:] 是命令行参数
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d] = %v\n", i, v)
	}
	fmt.Println("a = ", a)
}
