package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子，只需要一次
	rand.Seed(time.Now().UnixNano())
	// 产生随机数
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println("生成的随机数为：", array)
	// 冒泡排序
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("冒泡排序后：", array)

	// 数组做函数参数时，是值传递，不会改变原数组的值
	// 切片
	slice := array[1:5:5] // [1, 5) 5是容量，容量是可选的
	fmt.Println("slice = ", slice)
}
