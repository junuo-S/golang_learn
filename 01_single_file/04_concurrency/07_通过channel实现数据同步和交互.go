package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)
	go func() {
		defer fmt.Println("defer in sub goroutine")
		for i := 0; i < 2; i++ {
			fmt.Println("sub i = ", i)
			time.Sleep(time.Second)
		}
		ch <- 1 // 无缓存的 channel，发送方会阻塞，直到接收方接收数据
	}()
	<-ch
	close(ch) // 可以主动关闭channel
	//num, ok := <-ch // 如果ok为false，说明channel已经关闭
	// 也可以通过range遍历channel内容
}
