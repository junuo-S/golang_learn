package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，设置时间为2s，2s后，向time通道写数据（当前时间）
	timer := time.NewTimer(2 * time.Second) // 时间到了只会响应一次
	fmt.Printf("当前时间：%v\n", time.Now())
	fmt.Printf("2s后的时间：%v\n", <-timer.C) // 2s后，往timer.C写数据，有数据后，读取数据。如果没有数据，会阻塞

	timer.Stop()                 // 可以停止定时器
	timer.Reset(3 * time.Second) // 重置定时器
	// 通过timer实现延时功能
	// 延时2s后，打印一句话
	<-time.After(2 * time.Second)
	fmt.Println("2s后的时间：", time.Now())
}
