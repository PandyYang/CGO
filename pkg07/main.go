package main

//#include "ntisdk.h"
//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

import "fmt"

func main() {
	fmt.Println(C.query(C.CString("123")))
}
