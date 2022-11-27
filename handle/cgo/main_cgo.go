package main

// #include "libsdk.h"
// #cgo LDFLAGS: -L. -lsdk
import "C"

func main() {
	context := C.initConf(C.CString("/root/goProject/src/CGO/handle/cgo/conf.ini"))
	res := C.query(context, C.CString("123"))
	println(C.GoString(res))
}
