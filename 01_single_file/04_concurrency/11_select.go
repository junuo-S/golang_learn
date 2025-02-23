package main

import "fmt"

func main() {
	ch := make(chan int)    // 数字通信
	quit := make(chan bool) // 程序是否结束

	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true
	}()

	// 消费者，从ch读取内容

	// 生产者，生成数字，写入ch
	fibonacci(ch, quit)
}

// ch只写，quit只读
func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		// 监听channel数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}
