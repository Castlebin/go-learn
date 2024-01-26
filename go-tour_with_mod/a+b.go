package main

import "fmt"

/*
go 语言的 a + b  Online Judge
*/
func main() {
	var a, b int
	fmt.Scan(&a, &b) // 从标准输入中读取两个整数，赋值给 a 和 b . fmt.Scan() 入参是可变参数，可以接受任意个参数。

	sum := a + b
	fmt.Println(sum)
}
