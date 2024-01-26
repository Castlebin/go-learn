package main

import (
	"fmt"
	"github.com/timandy/routine"
	"strconv"
	"time"
)

// goroutine 是由 Go 运行时管理的轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
// 也就是我们说的 协程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

	// 随便看下 goroutine id
	fmt.Println("go say: " + strconv.Itoa(int(routine.Goid())))
}

func main() {
	go say("world")
	say("hello")

	// 随便看下 goroutine id
	fmt.Println("main: " + strconv.Itoa(int(routine.Goid())))
}
