package main

import "fmt"

func main() {
	var flag = true
	fmt.Println("flag = ", flag) // true

	// bool类型不能转换为int
	// fmt.Printf("flag = %d\n", int(flag)) // true

	// int类型不能转换为bool, 这种不能转换的类型，叫做不兼容类型
	// flag = bool(0) // cannot convert 0 (untyped int constant) to bool

	var ch byte
	ch = 'a' // 字符类型本质上是整型
	var t int
	// t = ch // cannot use ch (variable of type byte) as int value in assignment
	t = int(ch)            // 字符类型是可以直接赋值给整型的, 类型转换
	fmt.Println("t = ", t) // 97
}
