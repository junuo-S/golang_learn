package main

import "fmt"

type Hummaner interface {
	// 方法声明，只有声明，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

type Personer interface {
	Hummaner // 匿名字段，继承了Hummaner接口
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (s *Student) sayHi() {
	fmt.Printf("Student[%s, %d] sayHi\n", s.name, s.id)
}

func (s *Student) sing(lrc string) {
	fmt.Printf("Student[%s, %d] sing[%s]\n", s.name, s.id, lrc)
}

type Teacher struct {
	addr  string
	group string
}

func (t *Teacher) sayHi() {
	fmt.Printf("Teacher[%s, %s] sayHi\n", t.addr, t.group)
}

// 定义一个普通函数，函数的参数为接口类型
func WhoSayHi(i Hummaner) {
	i.sayHi()
}

func main() {
	// 定义一个接口类型的变量
	var i Hummaner = &Student{"mike", 666}
	i.sayHi()
	var inf = &Teacher{"bj", "go"}
	inf.sayHi()
	WhoSayHi(&Student{"mike", 666})
	WhoSayHi(&Teacher{"bj", "go"})
	var p Personer = &Student{"mike", 666}
	p.sayHi()     // 继承的方法
	p.sing("好汉歌") // 自己的方法

	// 接口转换，超集可以转换为子集，反过来不行
	i = p
	i.sayHi()

	// p = i // error

	// 空接口，万能类型，可以保存任意类型的值，类似于C语言的void*
	var t interface{} = 1
	fmt.Println("t = ", t)
	t = "abc"
	fmt.Println("t = ", t)

	// 类型断言
	var x interface{}
	x = "hello mike"
	v, ok := x.(string)
	fmt.Println("v = ", v, "ok = ", ok)

	switch t := x.(type) {
	case string:
		fmt.Println("x is string", t)
	case int:
		fmt.Println("x is int", t)
	case byte:
		fmt.Println("x is byte", t)
	}
}
