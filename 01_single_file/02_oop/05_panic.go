package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaa")
}

func testb() {
	fmt.Println("bbbbbbbb")
	panic("this is a panic testb") // 显式调用 panic 函数，导致程序崩溃
}

func testc() {
	fmt.Println("cccccccc")
}

// 数组越界
func testd() {
	var a [10]int
	for i := 0; i <= 10; i++ {
		a[i] = i
	}
}

// recover
func teste() {
	defer func() {
		if err := recover(); err != nil { // 产生了 panic 异常
			fmt.Println("recover from", err)
		}
	}()
	panic("this is a panic testd")
}

func main() {
	testa()
	//testb()
	testc()
	//testd()
	teste()
}
