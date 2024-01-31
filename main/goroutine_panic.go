package main

import (
	"strconv"
	"time"
)

// goroutine panic 会导致整个进程退出
func main() {
	go func() {
		for i := 0; i < 10; i++ {
			println("n: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	go func() {
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
