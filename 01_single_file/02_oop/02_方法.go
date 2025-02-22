package main

import "fmt"

type person struct {
	name string
	age  int
	sex  byte
}

// 带有接收者的函数叫方法
func (this person) PrintInfo() {
	fmt.Printf("%+v\n", this)
}

// 通过一个函数，给成员赋值
func (p *person) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

// 方法继承
type student struct {
	person
}

// 重写方法
func (this student) PrintInfo() {
	fmt.Printf("student: %+v\n", this)
}

func main() {
	var per = person{"mike", 18, 'm'}
	per.PrintInfo()
	per.SetInfo("yoyo", 20)
	per.PrintInfo()

	var stu = student{person{"yoyo", 20, 'f'}}
	stu.PrintInfo() // 继承了person的方法
	// 显式调用person的方法
	stu.person.PrintInfo()
}
