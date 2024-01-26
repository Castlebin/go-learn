package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 一般情况下，channel 无需关闭。只有在需要告诉接收方 goroutine 已经没有需要发送的值的时候才有必要关闭。
	// 例如终止一个 range 循环。
	// 只有发送者才有必要关闭 channel，接收者不应该去关闭 channel。
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

	// 可以用 v, ok := <-ch 来判断 channel 是否被关闭。
	// 如果 ok 是 false，那么说明 channel 已经没有任何数据并且已经被关闭。
	// 重要提示：只有在知道对方 goroutine 会一直向 channel 发送数据，而不会关闭 channel 的情况下才使用这种方式。
	v, ok := <-c
	if !ok {
		fmt.Println("channel is closed")
	} else {
		fmt.Println(v)
	}
}
