package main

import "fmt"

type IAb interface {
	Ab()
}

type A1 struct{}

func (a A1) Ab() {}

type B1 struct{}

// 这里的 Ab() 方法是指针接收者，而不是值接收者。
// 也就是说，B1 没有实现 IAb 接口。
// go 语言，当前并不认为接受者是指针时，算是实现了接口的方法。只能说，呵呵哒
func (b *B1) Ab() {}

func main() {
	var v IAb

	a1 := A1{}
	b1 := B1{}

	v = a1

	// v = b1      // 编译报错，说 B1 没有实现 IAb 接口，因为 B1 的 Ab() 方法是指针接收者，而不是值接收者。

	// 仅仅是为了编译通过，防止编译器说变量声明了，但是没有使用。
	fmt.Println(b1)
	fmt.Println(v)
}
