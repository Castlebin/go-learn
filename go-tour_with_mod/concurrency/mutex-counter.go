package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.Mutex 互斥锁 。相当于 java 中的 Lock。保证每次只有一个 Go 程能够访问 临界区   （需要互斥访问的代码块，通常是为了保证一些变量的状态在并发情况下的正确性）
*/

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()

	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++

	defer c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
