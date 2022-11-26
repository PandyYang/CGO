package main

import (
	"fmt"
)

//#cgo LDFLAGS:
//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
//char* echo(char* s);
import "C"

func main() {

}

//export query
func query(path string) string {
	fmt.Printf(path)
	return path
}
