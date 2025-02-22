package main

import "fmt"

func main() {
	a := 1
	// 匿名函数
	f1 := func() {
		fmt.Println("f1", a) // 捕获外部变量，在里面改了值，外部也会改变。
		// 闭包以引用方式捕获外部变量。
		//捕获的变量无论是否超出作用域，只要闭包还在用都不会被释放。
	}
	f1()
	func(a, b int) int {
		return a + b
	}(1, 2) // 直接调用
}
