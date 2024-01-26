package main

/*
*
闭包：可以理解为带有状态的函数

闭包是一个函数值，它引用了其函数体之外的变量。

是一个 返回带有状态的函数 的函数
*/
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		tmp := a
		a, b = b, a+b
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		println(f())
	}
}
