package main

// #cgo LDFLAGS: -ldl
// 启用CGO的特性 gobuild会在编译和链接阶段启动gcc编译器

import "C"

func main() {
	println("hello cgo")
}
