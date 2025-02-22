package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("a = ", a) // 变量声明后没有赋值，系统会自动赋值为0
	a = 1
	b = 2
	var c int = a + b
	fmt.Println("c = ", c) // 变量定义后必须使用，否则会报错

	var d, e int = 3, 4
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

	var f = 5 // 自动推导类型
	fmt.Println("f = ", f)

	g := 6 // 自动推导类型
	fmt.Println("g = ", g)
}
