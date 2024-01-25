package main

import "fmt"

// 泛型 结构体声明
// List 代表了一个可以支持任意类型值的链表
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	list := List[int]{}
	list.val = 1

	list.next = &List[int]{}
	list.next.val = 2

	list.next.next = &List[int]{}
	list.next.next.val = 3

	fmt.Println(list)
}
