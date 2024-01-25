package main

/**
type switch 语法
*/

func main() {
	do(21)
	do("hello")
	do(true)
}

/*
*
type switch 语法
*/
func do(i interface{}) { // i 的类型是 interface{} , 是具体的类型的划，编译都会报错
	// type switch 也只能作用于 声明为 interface{} 类型的变量
	switch v := i.(type) { // 可以看到这里的 v 类型也是 interface{}
	case int:
		println("int", v)
	case string:
		println("string", v)
	default:
		println("unknown", v)
	}
}
