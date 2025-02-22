package main

import "fmt"

func main() {
	var a int = 10
	var p *int = &a
	fmt.Println("a = ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	p = new(int) // go 只管分配内存，不管释放，会自动GC
	*p = 100
	fmt.Println("p = ", p)
}
