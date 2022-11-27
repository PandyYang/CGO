package main

import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

var (
	globalConfig string
)

func main() {

}

//export query
func query(path *C.char) {
	fmt.Println("get config path %s", C.GoString(path))
}

//export start
func start() {
	// just keep running to keep globalConfig in RAM
	// go forever()
	// quitChannel := make(chan os.Signal, 1)
	// signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	// <-quitChannel
}

func forever() {
	v := "global value"
	globalConfig = v
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(globalConfig)
	}
}

var (
	defaultConfigInstance *viper.Viper
)

func LoadNewConfigFile(path string) *viper.Viper {

	if defaultConfigInstance != nil {
		return defaultConfigInstance
	}

	fmt.Println("init conf")

	var v = viper.New()
	v.SetConfigFile(path)
	v.ReadInConfig()
	defaultConfigInstance = v
	return defaultConfigInstance

	//if defaultConfigInstance == nil {
	//	defaultConfigInstance = NewConfig("")
	//}
	//onceConfig.Do(func() {
	//	defaultConfigInstance.SetConfigFile(path)
	//	defaultConfigInstance.MergeInConfig()
	//})
}
