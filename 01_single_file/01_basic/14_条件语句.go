package main

import "fmt"

func main() {
	var s = "hello"
	if s == "hello" {
		fmt.Printf("s is %s\n", s)
	}

	// if 支持初始化语句
	if a := 10; a == 10 {
		fmt.Printf("a is %d\n", a)
	} else {
		fmt.Printf("a is not %d\n", a)
	}
}
