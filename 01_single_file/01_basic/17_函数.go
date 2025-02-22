package main

import "fmt"

// 无参无返回值函数，函数定义放前面后面都可以
func f1() {
	println("f1")
}

// 有参无返回值函数
// 参数传递只是值传递，不会影响原始值
func f2(a int, b string) {
	println(a, b)
}

// 不定参数函数
// 不定参数函数的参数只能有一个，并且必须是最后一个参数
// 传递的参数可以是 0 个或多个
func f3(a int, b ...string) {
	fmt.Printf("%T\n", b) // []string
	fmt.Println(a, b)
	// f4(b...) // 传递不定参数
	f4(b[:2]...) // 传递切片 左闭右开
}
func f4(a ...string) {
	fmt.Println(a)
}

// 无参有一个返回值函数
func f5() int {
	return 1
}

// 给返回值命名，go推荐的写法
func f6() (result int) {
	result = 1 // 只需要给 result 赋值即可
	return
}

// 多个返回值函数
func f7() (int, string) {
	return 1, "hello"
}
func f8() (a int, b string) {
	a = 1
	b = "hello"
	return // 有返回值的函数必须有 return 语句
}

// 有参有返回值函数
func MyMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f1()
	f2(1, "hello")
	f3(1, "hello", "world")
	println(f5())
}
