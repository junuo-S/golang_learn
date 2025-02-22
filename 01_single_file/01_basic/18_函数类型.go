package main

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 函数类型
// 函数也是一种类型，可以通过 type 给一个函数类型起名
type opFunc func(int, int) int

func main() {
	var f opFunc
	f = add
	println(f(1, 2))
	f = sub
	println(f(1, 2))
}
