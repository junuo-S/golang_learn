package main

import "fmt"

// 定义一个结构体
type Student struct {
	id   int
	name string
	sex  byte // 字符类型
	age  int
	addr string
}

func main() {
	// 顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s1 = ", s1)

	// 部分成员初始化
	s2 := Student{addr: "sh"}
	fmt.Println("s2 = ", s2)

	// 指针变量
	var s3 *Student = &Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s3 = ", s3)

	// 成员操作
	s3.id = 2
	s3.name = "yoyo"
	fmt.Println("s3 = ", s3)

	// 结构体比较，只能比较相同类型的结构体，只支持 == 或 !=
	s4 := Student{1, "mike", 'm', 18, "bj"}
	s5 := Student{1, "mike", 'm', 18, "bj"}
	fmt.Println("s4 == s5: ", s4 == s5)

	// 结构体做函数参数，值传递，形参无法改变实参
}
