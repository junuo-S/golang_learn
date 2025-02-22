package main

import "fmt"

func main() {
	a := 10
	b := "hello"
	c := 3.14
	d := true
	f := 1 + 2i
	g := 'a'
	fmt.Printf("a type: %T, value: %d\n", a, a)
	fmt.Printf("b type: %T, value: %s\n", b, b)
	fmt.Printf("c type: %T, value: %f\n", c, c)
	fmt.Printf("d type: %T, value: %t\n", d, d)
	fmt.Printf("f type: %T, value: %v\n", f, f)
	fmt.Printf("g type: %T, value: %c\n", g, g)

	// %d 整数
	// %s 字符串
	// %f 浮点数
	// %t 布尔
	// %v 自动匹配
	// %c 字符
}
