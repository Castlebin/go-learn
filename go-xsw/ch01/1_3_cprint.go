package main

// 运行报错：go: no Go source files。暂时还不知道为啥!
/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("Hello, world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
