package main

import "fmt"

// Index 方法返回 x 在 s 中的索引位置，如果 s 中不存在 x 则返回 -1。
// Index 方法的类型参数 T 有一个约束条件，即 T 必须是 可比较的类型 [T comparable]。
// 由于 T 是可比较的，所以我们可以在 Index 方法中使用 == 比较 x 和 v
// 函数的泛型参数，在函数名后面，参数列表前面，使用 [T XxType] 的形式声明
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x { // 实现了
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
