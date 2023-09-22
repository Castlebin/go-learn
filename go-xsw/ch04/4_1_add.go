package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

// 主 goroutine (main) 结束。程序并不会等待其他的 goroutine 结束。所以这段代码，很多时候都有可能是看不到任何输出的
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
