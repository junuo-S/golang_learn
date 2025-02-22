package main

func main() {
	// for 循环
	for i := 0; i < 10; i++ {
		println(i)
	}

	// for 循环的初始化语句和后置语句是可选的
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// 无限循环
	// for {
	// 	println("loop")
	// }

	// for range 循环
	s := "hello"
	// range 返回两个值，第一个是索引，第二个是值
	for i, v := range s {
		println(i, v)
	}

	// 第二个返回值默认丢弃，可以使用 _ 忽略
	for i := range s {
		println(i)
	}
}
