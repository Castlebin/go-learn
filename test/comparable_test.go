package test

import (
	"errors"
	"fmt"
	"testing"
)

/*
*
虽然 go 有对 comparable 的语义定义，但是一般来说，建议还是不要直接使用  `==` 或者 `!=` 来直接比较两个对象

对于自定义的结构体类型，实现一个 `Equals` 方法，是一个比较好的选择。然后用自定义的这个 `Equals` 方法去做比较
*/
func Test_Error_Comparable(t *testing.T) {
	// 不要试图用 `==` 或者 `!=` 去比较 error 类型
	err1 := errors.New("错误信息！")
	err2 := errors.New("错误信息！")
	fmt.Println("err1 == err2 ? ", err1 == err2) // false
	// fmt.Println("err1 == err2 ? ", *err1 == *err2) // 虽然它是指针，但是没法对它做 * 操作，所以，不要试图用 `==` 或者 `!=` 去比较 error 类型

	fmt.Println("err1.Error() == err2.Error() ? ", err1.Error() == err2.Error()) // true

	t1 := Test_Comparable{A: "1", B: 2}
	t2 := Test_Comparable{A: "1", B: 2}
	fmt.Println("t1 == t2 ? ", t1 == t2) // true

	p1 := &t1
	p2 := &t2
	fmt.Println("p1 == p2 ? ", p1 == p2)     // of course it's false
	fmt.Println("*p1 == *p2 ? ", *p1 == *p2) // true，因为这就是  t1 == t2
}

type Test_Comparable struct {
	A string
	B int
}
