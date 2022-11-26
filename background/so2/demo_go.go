package main

// #include "libtest.h"
// #cgo LDFLAGS:-L ./ -ltest
import "C"
import "fmt"

func main() {
	fmt.Println(C.Add(1, 2))
}
