package main

/*
#include<stdio.h>
#include <stdlib.h>
#include <dlfcn.h>
static void SayHello(const char* s) {
	puts(s)
}
*/

// void SayHello2(const char* s);
import "C"

func main() {
	C.SayHello(C.CString("hello world\n"))
	C.SayHello2(C.CString("Hello, World123\n"))
}
