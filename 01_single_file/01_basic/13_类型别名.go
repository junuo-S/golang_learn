package main

import "fmt"

func main() {
	type MyInt int
	type MyFloat float64
	type MyBool bool

	var a MyInt = 1
	var b MyFloat = 1.1
	var c MyBool = true

	fmt.Printf("a type is: %T\n", a)
	fmt.Printf("b type is: %T\n", b)
	fmt.Printf("c type is: %T\n", c)

	type (
		MyInt2   int
		MyFloat2 float64
		MyBool2  bool
	)
	fmt.Println(10 / 3) // 整除
}
