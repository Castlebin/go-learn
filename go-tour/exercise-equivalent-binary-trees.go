package main

// 练习：等价二叉查找树
import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// 遍历结束，关闭 channel ，这样可以提示接收端 ，遍历已经结束
func Walking(t *Tree, ch chan int) {
	Walk(t, ch)
	defer close(ch) // close channel after Walk() finishes
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walking(t1, ch1)
	go Walking(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(&Tree{Value: 1}, &Tree{Value: 1}))
	fmt.Println(Same(&Tree{Value: 1}, &Tree{Value: 2}))
}
