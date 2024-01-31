package main

import (
	"strconv"
	"time"
)

// goroutine panic 会导致整个进程退出         (java 中的线程、虚拟线程都不会)
// 要在 go routine 中捕获 panic，需要使用 recover。采用 defer + recover 的方式，可以捕获到 panic，解决该问题。
// 但是，必须要在 goroutine 内部 recover，才能捕获到 panic。如果在 go routine 外部 recover，是捕获不到这个 panic 的。
func main() {
	go func() {
		for i := 0; i < 10; i++ {
			println("n: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				println("panic: " + err.(string))
			}
		}()

		for i := 0; i < 10; i++ {
			println("p: " + strconv.Itoa(i))
			time.Sleep(time.Second)
			if i == 3 {
				panic("test panic")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		println("m: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
