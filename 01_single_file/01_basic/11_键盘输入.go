package main

import "fmt"

func main() {
	var name string
	fmt.Printf("请输入你的名字：")
	fmt.Scanf("%s", &name)
	fmt.Printf("你好，%s\n", name)
	fmt.Scan(&name)
	fmt.Printf("你好，%s\n", name)
}
