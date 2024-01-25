package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * 练习：实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。

 io.Reader 接口声明如下：

type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

type MyReader struct{}

/*
*

	注意，这里的 Read() 方法是值接收者，而不是指针接收者。

为什么呢？因为，go 语言认为，只有值接收者，才算是实现了接口的方法。
*/
func (myReader MyReader) Read(b []byte) (n int, err error) {
	n = len(b)
	for i := 0; i < n; i++ {
		b[i] = 'A'
	}
	return n, nil
}

func main() {
	Validate(MyReader{})
}

func Validate(r io.Reader) {
	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ { // test 1mb
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				fmt.Fprintf(os.Stderr, "got byte %x at offset %v, want 'A'\n", v, o+i)
				return
			}
		}
		o += n
		if err != nil {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
			return
		}
	}
	if o == 0 {
		fmt.Fprintf(os.Stderr, "read zero bytes after %d Read calls\n", i)
		return
	}
	fmt.Println("OK!")
}
