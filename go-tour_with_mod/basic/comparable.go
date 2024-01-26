package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"Alice", 30}
	p2 := Person{"Bob", 25}

	fmt.Println("Are p1 and p2 equal?", p1 == p2)     // 使用 == 进行相等性判断
	fmt.Println("Are p1 and p2 not equal?", p1 != p2) // 使用 == 进行相等性判断

	fmt.Println("Is p1 older than p2?", p1.Age > p2.Age) // 使用 > 进行大小比较
}

/**
Go 语言中，类型 comparable 的定义如下：   （comparable 的类型的变量，可以使用 == 和 != 折两个比较操作。不能使用大于小于这些）

在 Go 语言中，你无法自定义一个类型是否是可比较的，可比较性是由编译器根据类型的属性来确定的。只有满足特定条件的类型才能被视为可比较的。

Go 语言规范中规定了可比较类型的属性，包括：

基础类型（如数值类型、字符串类型、布尔类型）
指针类型
数组类型
结构体类型（当结构体中的所有字段都是可比较类型时）
如果你想要自定义一个类型，并希望它是可比较的，你需要确保它的属性满足上述规范。如果你需要自定义比较行为，可以为该类型定义方法来实现自定义的比较操作，例如实现 Equals 方法来进行自定义的相等性判断。

如果你有特定的需求，可以考虑使用接口来定义自定义的比较行为，然后为不同的类型实现这个接口。这样可以让你在不同的类型上实现相同的比较逻辑。


*/
