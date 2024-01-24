package test

import (
	"fmt"
	"testing"
)

/*
在 go 中，获取变量的地址，可以使用 & 符号：
打印变量 a 的地址，可以使用 fmt.Printf("%p", &a)。
*/
func Test_var_addr(t *testing.T) {
	// 打印 go 中的变量的地址的步骤：
	// 1. 使用 & 符号，获取变量的内存地址
	// 2. 使用 fmt.Printf("%p", &a) 打印变量的地址
	a := 1
	fmt.Printf("变量 a 的内存地址是：%p\n", &a)

	var s string
	fmt.Printf("字符串变量 s 的内存地址是：%p\n", &s)

	m := make(map[string]int)
	fmt.Printf("map 变量 m 的内存地址是：%p\n", &m)

	var aa AA
	fmt.Printf("结构体变量 aa 的内存地址是：%p\n", &aa)
}

type AA struct{}
