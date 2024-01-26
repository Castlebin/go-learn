package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// 通道（channel）是用来传递数据的一个数据结构。
// channel 空的话，接收方会阻塞，直到有数据发送过来
// channel 满的话，发送方会阻塞，直到数据被接收
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
