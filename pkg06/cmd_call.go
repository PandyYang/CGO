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
	"os"
	"path/filepath"
)

type FileInfo struct {
	FileName       string `json:"file_name,omitempty"`
	Md5Checksum    string `json:"md5_checksum,omitempty"`
	Sha1Checksum   string `json:"sha1_checksum,omitempty"`
	Sha256Checksum string `json:"sha256_checksum,omitempty"`
	Size           uint64 `json:"size,omitempty"`
	LocalPath      string `json:"local_path,omitempty"`
}

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
	ReadFromCMD("E:\\GolandProject\\CGO\\pkg06\\123.txt")
}

//export ReadFromCMD
func ReadFromCMD(path string) {
	// fname := C.CString(path)
	// defer C.free(unsafe.Pointer(fname))
	fileBytes, err := os.ReadFile(fmt.Sprintf("%s", path))
	fmt.Println(path, "=============================11111111111111")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(string(fileBytes))
	fmt.Println(path, "=============================222222222222222")

}

func single(fFile string) int {
	fInfo, err := os.Stat(fFile)
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
	fileInfo := new(FileInfo)
	fileInfo.LocalPath = v
	err = readFile(fileInfo)
	if err != nil {
		log.Println("read file error", err.Error())
		return -1
	}

	return 0
}

// 解析文件内容
func readFile(file *FileInfo) error {
	path := file.LocalPath
	fileBytes, err := os.ReadFile(path)

	if err != nil {
		log.Println("when read file, something error", err.Error())
		return err
	}

	fmt.Println(string(fileBytes))
	return nil
}
