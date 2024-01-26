package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// hello 和 world 会交替输出
	go say("world") // 很有可能看到输出的 world 少于 5 次    （因为 go 的主线程 结束了，go routinues 也会结束）
	say("hello")
}
