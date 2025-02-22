package main

import "fmt"

func test(x int) {
	fmt.Println(1 / x)
}

func main() {
	fmt.Println("start")
	// defer 语句即使运行时出错，也会执行
	defer fmt.Println("defer1") // defer 语句会在函数执行完毕后执行，类似于 finally
	defer fmt.Println("defer2") // defer 语句会按照先进后出的顺序执行 FILO
	fmt.Println("center")

	test(0)                     // 除零错误
	defer fmt.Println("defer3") // 不会执行，因为已经 panic 了
}
