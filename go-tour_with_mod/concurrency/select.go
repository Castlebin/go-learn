package main

import "fmt"

/*
select 语句:
select 语句使一个 Go 程可以等待多个通信操作。
select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

这个特性可以用在以下场景：
1. 可以从多个数据源接收数据，接收到一个就行。其他的就不用等了
2. 一个 channel 用来接收数据，一个 channel 用来发送数据，这样就可以实现超时机制。或者可以用来实现被动取消机制
3.。。。
*/
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // c channel 可以发送数据
			x, y = y, x+y
		case <-quit: // quit channel 可以接收数据
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
