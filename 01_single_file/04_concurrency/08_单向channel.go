package main

func main() {
	// 创建一个channel
	ch := make(chan int) // 默认是双向的
	// 双向channel 能隐式转换为单向channel
	var writeCh chan<- int = ch // 只能写不能读
	writeCh <- 666              // 写
	//<- writeCh // invalid operation: <-writeCh (receive from send-only type chan<- int)
	var readCh <-chan int = ch // 只能读不能写
	<-readCh                   // 读
	// readCh <- 666 // invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	// 单向channel无法转换为双向
	// var ch2 chan int = writeCh // cannot use writeCh (type chan<- int) as type chan int in assignment
	// var ch3 chan int = readCh // cannot use readCh (type <-chan int) as type chan int in assignment

	// 单向channel的应用
	ch2 := make(chan int)
	// 生产者，生产数字，写入channel
	go producer(ch2) // channel传参，引用传递
	// 消费者，从channel读取数字，打印
	consumer(ch2)
}

func consumer(ch <-chan int) {
	for num := range ch {
		println(num)
	}
}

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
