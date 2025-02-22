package main

import "fmt"

func main() {
	// iota 常量自动生成器，每个一行，自动累加1
	const (
		a = iota // iota是go语言的常量计数器，只能在常量的表达式中使用
		b = iota
		c = iota
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	// iota遇到const，重置为0
	const d = iota
	fmt.Printf("d = %d\n", d)

	// 可以只写一个iota
	const (
		e = iota
		f
		g
	)
	fmt.Println("e = ", e, ", f = ", f, ", g = ", g)

	// 如果是同一行，值都一样
	const (
		h, i, j = iota, iota, iota
	)
	fmt.Println("h = ", h, ", i = ", i, ", j = ", j)
}
