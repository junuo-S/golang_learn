package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 1. 解释规则
	// 2. 根据规则提取关键信息

	buf := "abc azc a7c aac 888 a9c tac"
	// 解释规则, 它会解析正则表达式，如果成功返回解释器
	re := regexp.MustCompile(`a.c`)
	if re == nil { // 解释失败，返回nil
		fmt.Println("MustCompile err")
		return
	}
	// 根据规则提取关键信息
	result1 := re.FindAllStringSubmatch(buf, -1) // -1表示返回所有匹配的个数
	fmt.Println("result1 = ", result1)

	re2 := regexp.MustCompile(`a[0-9]c`) // `a\dc`
	if re2 == nil {
		fmt.Println("MustCompile err")
		return
	}
	result2 := re2.FindAllStringSubmatch(buf, -1)
	fmt.Println("result2 = ", result2)

	// 提取有效的小数
	buf = "3.14 567 agsdg 1.23 7. 8.99 1. 123 7.8"
	re3 := regexp.MustCompile(`\d+\.\d+`) // +表示匹配前面的字符1次或多次
	if re3 == nil {
		fmt.Println("MustCompile err")
		return
	}
	result3 := re3.FindAllStringSubmatch(buf, -1)
	fmt.Println("result3 = ", result3)
}
