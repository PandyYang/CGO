package main

import "C"
import "log"

//export RunLib
func RunLib() {
	log.Println("Call RunLib")
}

func init() {
	log.Println("Call init")
}

func main() {
	log.Println("Call main")
}
