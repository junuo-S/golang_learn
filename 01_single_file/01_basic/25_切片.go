package main

import "fmt"

func main() {
	// 生成一些切片代码
	sl := make([]int, 5, 10) // cap可以省略，默认和len一样
	fmt.Println("len(sl) = ", len(sl), "cap(sl) = ", cap(sl))

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// [low, high, max] low是切片的开始索引，high是切片的结束索引. len = high - low, cap = max - low
	slice := array[1:5:5]
	fmt.Println("slice = ", slice, "len(slice) = ", len(slice), "cap(slice) = ", cap(slice))
	// 操作切片的某个元素 和 数组操作方式一样
}
