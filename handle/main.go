package main

/*
   extern void MyGoPrint(void *context);
   static inline void myprint(void *context) {
       MyGoPrint(context);
   }
*/
import "C"
import (
	"fmt"
	"runtime/cgo"
	"unsafe"

	"github.com/spf13/viper"
)

type Config struct {
	Name           string
	Age            int
	configInstance *viper.Viper
}

//export MyGoPrint
func MyGoPrint(context unsafe.Pointer) {
	h := *(*cgo.Handle)(context)
	val := h.Value().(Config)
	println(val.Name)
	println(val.Age)
	//h.Delete()
}

//export initConf
func initConf(filePath *C.char) unsafe.Pointer {
	cc := LoadNewConfigFile(C.GoString(filePath))
	c := Config{}
	c.Age = 20
	c.Name = "pandy"
	c.configInstance = cc
	val := c
	h := cgo.NewHandle(val)
	return unsafe.Pointer(&h)
}

//export query
func query(context unsafe.Pointer, str *C.char) *C.char {
	h := *(*cgo.Handle)(context)
	val := h.Value().(Config)
	res := val.configInstance.GetString("test.123")
	fmt.Println(res)
	return str
}

//export destroy
func destroy(context unsafe.Pointer) {
	h := *(*cgo.Handle)(context)
	h.Delete()
}

func main() {
	c := Config{}
	c.Age = 20
	c.Name = "pandy"

	val := c
	h := cgo.NewHandle(val)
	for i := 0; i < 1000000; i++ {
		C.myprint(unsafe.Pointer(&h))
		// Output: hello Go
	}

	h.Delete()

}

func LoadNewConfigFile(path string) *viper.Viper {
	var v = viper.New()
	v.SetConfigFile(path)
	v.ReadInConfig()
	return v
}
