package main

// #cgo LDFLAGS: -ldl
//#include<stdio.h>
import "C"

func main() {
	C.puts(C.CString("Hello, World\n"))
}
