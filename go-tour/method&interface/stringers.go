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

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
