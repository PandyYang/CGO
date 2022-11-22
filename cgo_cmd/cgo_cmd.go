package main

/*
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
*/
import "C"
import (
	"fmt"
	"log"
	"ntisdk/entries"
	"os"
	"path/filepath"
)

// 从命令行解析文件 填入文件路径 绝对相对皆可
// idea goland 打开edite configuration, 在Program arguments写入 -F xx/xx/xx.txt
func main() {
	// fmt.Println("start to read file")
	// fFile := pflag.StringP("file", "F", "", "file to read")
	// pflag.Parse()

	// if len(*fFile) > 0 {
	// 	single(*fFile)
	// } else {
	// 	fmt.Println("no file to read")
	// }
}

//export ReadFromCMD
func ReadFromCMD(path *C.char) {
	fmt.Printf("path type:%T\n", path)
	p := fmt.Sprintf("%s", C.GoString(path))
	// fname := C.CString(path)
	// defer C.free(unsafe.Pointer(fname))
	fmt.Println(p, "========================")
	i := single(p)
	if i == 5 {
		fmt.Println("can't read dir")
	} else if i == 255 {
		fmt.Println("open file failed")
	} else if i == -1 {
		fmt.Printf("read file failed")
	}
}

func single(fFile string) int {
	fmt.Println(fFile, "================")
	f := fFile
	fInfo, err := os.Stat(f)
	if err != nil {
		fmt.Println(err.Error())

		return 255
	}

	if fInfo.IsDir() {
		// Directory
		fmt.Printf("<%s> is a directory, use --dir or -D instead\n", fFile)
		return 5
	}

	// Original file
	///*
	v, _ := filepath.Abs(fFile)
	fileInfo := new(entries.FileInfo)
	fileInfo.LocalPath = v
	err = readFile(fileInfo)
	if err != nil {
		log.Println("read file error", err.Error())
		return -1
	}

	return 0
}

// 解析文件内容
func readFile(file *entries.FileInfo) error {
	path := file.LocalPath
	fileBytes, err := os.ReadFile(path)

	if err != nil {
		log.Println("when read file, something error", err.Error())
		return err
	}

	fmt.Println(string(fileBytes))
	return nil
}
