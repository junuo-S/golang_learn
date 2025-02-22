package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

type Student struct {
	Person // 匿名字段
	id     int
	addr   string
}

func main() {
	// 初始化
	var s1 = Student{Person{"mike", 18, 'm'}, 1, "bj"}
	fmt.Println("s1 = ", s1)
	// %+v 显示更详细
	fmt.Printf("s1 = %+v\n", s1)

	// 指定成员初始化，没有初始化的成员，自动赋值为0
	s2 := Student{id: 1}
	fmt.Printf("s2 = %+v\n", s2)
	s3 := Student{Person: Person{name: "mike"}}
	fmt.Printf("s3 = %+v\n", s3)

	// 访问成员
	s1.name = "yoyo"
	s1.age = 20
}
