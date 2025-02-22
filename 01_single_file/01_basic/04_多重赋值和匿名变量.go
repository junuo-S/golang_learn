package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)

	a, b = b, a // 交换两个变量的值
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)

	a, _, c = 4, 5, 6 // _是匿名变量，丢弃数据不处理，配合函数返回值使用
	fmt.Println("a = ", a, ", b = ", b, ", c = ", c)
}
