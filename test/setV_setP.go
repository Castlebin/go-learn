package main

import "fmt"

func main() {
	a := A{B{C{1}}}
	setV(a, 2)                           // setV inner:  2
	fmt.Println("after setV: ", a.B.C.C) // setV outer:  1   ，可以看到，setV 内部的赋值操作，没有影响到外层的 a

	setP(&a, 2)                          // setP inner:  2
	fmt.Println("after setP: ", a.B.C.C) // setP outer:  2 ，可以看到，C 值变了
}

type A struct {
	B B
}
type B struct {
	C C
}
type C struct {
	C int
}

// 函数声明 a 的类型为结构体本身
// （调用时，在该函数内部，是一个完全新的内存空间，深拷贝了一份，
// 对 a 的属性的任何赋值操作，都不会影响到调用该方法时，外层代码传入的 a 的属性）
func setV(a A, c int) {
	a.B.C.C = c
	fmt.Println("setV inner: ", a.B.C.C)
}

// 函数声明 a 为指针 （调用时，在该函数内部，是一个指针，对 a 的属性的任何赋值操作，都是直接操作调用时的原来的 a）
func setP(a *A, c int) {
	a.B.C.C = c
	fmt.Println("setP inner: ", a.B.C.C)
}

/**
对方法调用时，浅拷贝 和 深拷贝 的解释 （看类型） ：
 1、浅拷贝
  浅拷贝是指对地址的拷贝
  浅拷贝的是数据地址，只复制指向的对象的指针，
此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化 (其实没有新老对象，就是同一个)。

  引用类型的都是浅拷贝：slice、map、function   （+ 指针）    （注意：slice 在这里，和 array 是不一样的）

 2、深拷贝
  深拷贝是指将地址指向的值进行拷贝

  深拷贝的是数据本身，创造一个一样的新对象，新创建的对象与原对象不共享内存，
新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。

  值类似的都是深拷贝：int、float、bool、array、struct  （看到没有，struct 、array 都在这里。）
*/
