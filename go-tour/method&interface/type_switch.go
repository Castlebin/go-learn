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
	// 可以发现，go 的 type switch 语法，在 case 中，会进行类型转换，所以才能打印出来正确的值。如果我们这里加上 case bool，值也能打印出来。
	// 但是，如果执行 default 时，是没有做类型转换的，所以，default 中打印 bool 类型的值，不能打印出来 true/false ，而是两个 地址  (不知道什么意思)
	//case bool:
	//	println("bool", v)
	default:
		println("unknown", v)
	}
}
