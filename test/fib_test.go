package test

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	fmt.Println(fibs(10))
}

func fibs(n int) []int {
	var fibs []int
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fibs = append(fibs, b)
		a, b = b, a+b
	}
	return fibs
}
