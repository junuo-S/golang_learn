package main

import "fmt"

func main() {
	a := 1
	fmt.Println("a = ", a)    // 自动换行
	fmt.Printf("a = %d\n", a) // 格式化输出

	b := 2
	c := 3
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
