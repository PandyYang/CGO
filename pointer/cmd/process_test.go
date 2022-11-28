// package main

// import (
// 	"fmt"
// 	"ntisdk/utils"
// 	"testing"

// 	"github.com/panjf2000/ants/v2"
// )

// func TestStartProcess(t *testing.T) {
// 	startProcess("/root/goProject/src/NTI-SDK/conf/config.ini")
// }

// func TestQuery(t *testing.T) {

// 	ress := make(chan []string)
// 	utils.LoadConfigFile("/root/goProject/src/NTI-SDK/conf/config.ini")
// 	thread, _ := ants.NewPool(100)

// 	for i := 0; i < 10; i++ {

// 		thread.Submit(func() {
// 			res := search("0---00055.0000hugetits.0-108.btcc.com")
// 			ress <- res
// 		})
// 		fmt.Println(i)
// 		// fmt.Println(res)
// 	}

// 	for i := range ress {
// 		for _, v := range i {
// 			fmt.Println(v)
// 		}
// 	}

// }
