package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Sleep 1 second")
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Sleep 3 second")
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Sleep 5 second")
	}()

	// 只会输出这里的内容 ，因为主 goroutine 退出了，其他 goroutine 也会退出
	fmt.Println("main function")
}
