package main

/*
export LD_LIBRARY_PATH="xxx sopath xxx"
*/

// #include "libsdk.h"
// #cgo LDFLAGS: -L. -lsdk
import "C"

func main() {
	// C.start()
	for i := 0; i < 10; i++ {
		C.search()
	}
}
