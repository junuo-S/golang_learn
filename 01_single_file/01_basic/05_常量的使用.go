package main

import "fmt"

func main() {
	const name = "Golang"
	const age = 10
	// name = "Go" // 常量不能修改
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Printf("name type is : %T\n", name)
	fmt.Printf("age type is : %T\n", age)
}
