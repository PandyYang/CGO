package main

import "C"
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	globalConfig string
)

func main() {

}

//export search
func search() {
	fmt.Println(globalConfig)
}

//export start
func start() {
	// just keep running to keep globalConfig in RAM
	go forever()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
}

func forever() {
	v := "global value"
	globalConfig = v
	for {
		time.Sleep(10 * time.Second)
		fmt.Println(globalConfig)
	}
}
