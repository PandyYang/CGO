package main

// void SayHello2(const char* s);
import "C"

func main() {
	C.SayHello2(C.CString("Hello, World2222\n"))
}
