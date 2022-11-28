package main

// #include "libsdk.h"
// #cgo LDFLAGS:-L. -lsdk
import "C"
import "time"

func main() {

	sdkContext := C.initSDK(C.CString("/root/goProject/src/NTI-SDK/conf/config.ini"))
	// downloadFlag := C.downloadPackage(sdkContext)
	// if downloadFlag == 0 {

	for i := 1; i < 1000000; i++ {
		res := C.query(sdkContext, C.CString("123"))
		time.Sleep(100 * time.Microsecond)
		println(C.GoString(res))
		println(sdkContext)
		println(i)
	}

	// res := C.query(sdkContext, C.CString("0---00055.0000hugetits.0-108.btcc.com"))
	// println(C.GoString(res))
	// }

	C.destroy(sdkContext)
}

// type InitHandler struct {
// 	configInstance *viper.Viper
// 	redisPool      *redis.Pool
// 	loggerInstance *log.Logger
// }

// func parseContext(contex unsafe.Pointer) {
// 	h := *(*cgo.Handle)(context)
// 	handler := h.Value().(InitHandler)
// 	r := handler.configInstance.GetString("redis.PORT")
// 	println(r)
// }
