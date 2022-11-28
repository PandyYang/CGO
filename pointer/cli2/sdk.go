package main

// #include "libsdk2.h"
// #cgo LDFLAGS:-L. -lsdk2
import "C"
import "time"

func main() {

	sdkContext := C.initSDK(C.CString("/root/goProject/src/NTI-SDK/conf/config.ini"))
	// downloadFlag := C.downloadPackage(sdkContext)
	// if downloadFlag == 0 {

	for i := 1; i < 999; i++ {
		res := C.query(sdkContext, C.CString("321"))
		time.Sleep(100 * time.Microsecond)
		println(C.GoString(res))
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
