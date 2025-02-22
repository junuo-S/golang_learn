package main // go语言以包作为管理单位，每个文件必须先声明包，程序必须有一个main包

import "fmt" // 调用函数，大部分都需要导入对应的包

// 一个文件夹只能有一个main函数
func main() { // 左括号必须和函数同行
	fmt.Println("Hello, Golang!") // 此函数会自动换行, 语句末尾不需要加分号
}
