package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Append，转换为字符串后追加到字节数组
	s := make([]byte, 0, 100)
	s = strconv.AppendBool(s, true)
	s = strconv.AppendInt(s, 123, 10) // 第二个参数是要转换的数，第三个参数是进制
	s = strconv.AppendQuote(s, "abc")
	fmt.Println("s = ", string(s)) // 转换为字符串后输出

	// 其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)
	str = strconv.FormatFloat(3.14, 'f', -1, 64) // 'f' 表示格式，-1 表示精度，64 表示 float64
	fmt.Println("str = ", str)
	str = strconv.FormatInt(123, 10) // 10 表示进制
	fmt.Println("str = ", str)
	str = strconv.Itoa(666) // 常用
	fmt.Println("str = ", str)

	// 字符串转换为其他类型
	var flag, err = strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	// 把字符串转换为整型, 常用
	var num1, err1 = strconv.Atoi("666")
	if err1 == nil {
		fmt.Println("num1 = ", num1)
	} else {
		fmt.Println("err1 = ", err1)
	}
}
