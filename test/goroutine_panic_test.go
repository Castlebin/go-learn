package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// goroutine panic 会导致整个进程退出
func Test_Go_Routine_Panic(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			println("p: " + strconv.Itoa(i))
			time.Sleep(time.Second)
			if i == 3 {
				panic("test panic")
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			println("n: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		println("m: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func Test_Go_Routine_Panic_2(t *testing.T) {
	// 由于 panic 会导致整个进程退出，所以，这里的 panic 会导致整个进程退出。再深也没用，都是一样的
	go func() {
		go func() {
			go func() {
				for i := 0; i < 10; i++ {
					println("p: " + strconv.Itoa(i))
					time.Sleep(time.Second)
					if i == 3 {
						panic("test panic")
					}
				}
			}()
		}()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			println("n: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		println("m: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func Test_Go_Routine_Panic_Recover(t *testing.T) {
	go func() {
		go func() {
			go func() {
				// 得在 goroutine 内部 recover，才能捕获到 panic。采用 defer + recover 的方式
				defer func() {
					if r := recover(); r != nil {
						// 在这里处理panic
						fmt.Println("发生了panic:", r)
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
		}()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			println("n: " + strconv.Itoa(i))
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		println("m: " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
