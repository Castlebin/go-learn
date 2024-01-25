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

type AA struct {
	A int
}

/*
从 内存地址的打印结果，看方法传参中 值 和 指针的区别：

可以看到，当传递的是值的时候，方法内部的变量和外部的变量，是两个不同的变量，它们的内存地址是不同的。
当传递的是指针的时候，方法内部的变量和外部的变量，是同一个变量，它们的内存地址是相同的。

所以，指针传递时，能改变外部变量的值，值传递时，不能改变外部变量的值。
*/
func Test_var_addr2(t *testing.T) {
	a := 1
	b := AA{1}

	fmt.Printf("变量 a 的内存地址是：%p\n", &a)
	fmt.Printf("变量 b 的内存地址是：%p\n", &b)

	printAddr(a, &a, b, &b)

	fmt.Print(b)
}

func printAddr(a int, pa *int, b AA, pb *AA) {
	fmt.Printf("方法传参 a 的内存地址是：%p\n", &a)
	fmt.Printf("方法传参 指针 a 地址是：%p\n", pa)
	fmt.Printf("方法传参 b 的内存地址是：%p\n", &b)
	fmt.Printf("方法传参 指针 b 地址是：%p\n", pb)
}
