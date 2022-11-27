package main

// #include "libsdk.h"
// #cgo LDFLAGS: -L. -lsdk
import "C"
import "time"

func main() {
	// C.command(C.CString(""))
	for i := 0; i < 10; i++ {
		C.search(C.CString("123"))
	}

	time.Sleep(10 * time.Second)

	for i := 0; i < 10; i++ {
		C.search(C.CString("321"))
	}
}
