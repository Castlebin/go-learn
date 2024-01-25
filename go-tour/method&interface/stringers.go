package main

import "fmt"

/*
Stringers

fmt 包中定义的 Stringer 是使用最普遍的接口之一。

type Stringer interface {
	String() string
}

只要实现了 String 方法的类型，作为参数被 fmt.Println 打印时，都是会打印 String() 方法的值。
*/

type Person struct {
	Name string
	Age  int
}

// 这里的 String() 方法是值接收者  ， 只有是值接收者才会被 fmt.Println 调用。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// 会发现 打印 Person2 时，并没有调用 String() 方法，而是直接打印了 Person2 的值。
	b := Person2{"Arthur Dent", 42}
	y := Person2{"Zaphod Beeblebrox", 9001}
	fmt.Println(b, y)
}

type Person2 struct {
	Name string
	Age  int
}

// 区别在于，这里的 String() 方法是指针接收者，而上面的是值接收者。
// 这个跟 go 语言认为什么才算实现了接口有关。实现接口的方法，必须是值接收者。指针接收者，不算实现接口的方法。
func (p2 *Person2) String() string {
	return fmt.Sprintf("%v (%v years)", p2.Name, p2.Age)
}
