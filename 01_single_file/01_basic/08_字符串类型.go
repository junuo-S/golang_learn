package main

import "fmt"

func main() {
	// 字符串类型
	var str1 string = "hello"
	var str2 string = "world"
	var str3 string = str1 + str2
	fmt.Println("str3 = ", str3) // helloworld

	// 字符串长度
	fmt.Println("len(str3) = ", len(str3)) // 10

	// 字符串截取
	fmt.Println("str3[0:5] = ", str3[0:5])   // hello
	fmt.Println("str3[5:10] = ", str3[5:10]) // world

	// 字符串比较
	fmt.Println("str1 == str2: ", str1 == str2) // false
	fmt.Println("str1 != str2: ", str1 != str2) // true

	// 字符串遍历
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c ", str3[i])
	}
	fmt.Println()
	for i, v := range str3 {
		fmt.Printf("i = %d, v = %c\n", i, v)
	}
}
