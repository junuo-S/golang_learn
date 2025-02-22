package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5} // 定义的数组元素个数必须是常量, 可以部分赋值，没有赋值的默认为0
	fmt.Println(a)

	var p = new([5]int)
	for i := 0; i < 5; i++ {
		p[i] = i + 11
	}
	fmt.Println(p)

	var b = [5]int{2: 10} // 部分初始化，没有初始化的默认为0
	fmt.Println(b)

	// 支持两个类型相同的数组进行比较，只有当两个数组的长度和元素都相同时，才相等。只支持 == 和 !=
	// 同类型的数组可以赋值，但是长度不同的数组不能赋值
}
