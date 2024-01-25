package main

import "fmt"

/*
类型断言的语法

t := i.(T)                  // i 必须是接口类型，T 是普通类型
*/
func main() {
	// 类型断言，只能作用于 声明为 interface{} 类型的变量
	var t interface{} // t 的类型是 interface{}
	t = "hello"

	var x interface{} = "hello" // x 的类型是 interface{} , 如果这里把 interface{} 改为 string ，会报错。

	v, ok := t.(string)
	fmt.Println(v, ok)

	c, ok := x.(int)
	fmt.Println(c, ok)

}
