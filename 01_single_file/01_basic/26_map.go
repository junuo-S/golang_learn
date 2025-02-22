package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	// 对于map，只有len，没有cap
	m1 = make(map[int]string, 10) // 可以指定长度，但是只是指定了容量，不是长度，里面没有元素
	fmt.Println("m1 = ", m1)
	fmt.Println("len(m1) = ", len(m1))
	m1[1] = "mike" // 元素的操作，健是唯一的
	m1[2] = "go"
	m1[3] = "c++"
	fmt.Printf("m1 = %v\n", m1)
	fmt.Println("len(m1) = ", len(m1))

	// 定义并初始化
	m2 := map[int]string{1: "mike", 2: "go", 3: "c++"}
	fmt.Printf("m2 = %v\n", m2)

	// 遍历, 遍历的顺序是不确定的
	for k, v := range m2 {
		fmt.Printf("m2[%d] = %v\n", k, v)
	}

	// 判断某个键是否存在
	v, ok := m2[1]
	fmt.Printf("v = %v, ok = %v\n", v, ok)
	v, ok = m2[4]
	fmt.Printf("v = %v, ok = %v\n", v, ok)

	// 删除key = 1的元素
	delete(m2, 1)
	fmt.Printf("m2 = %v\n", m2)

	// map 做函数参数
	modify(m2) // 引用传递
	fmt.Printf("m2 = %v\n", m2)
}

func modify(m map[int]string) {
	m[1] = "mike~"
}
