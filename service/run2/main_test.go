package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestInitSDK(t *testing.T) {
	pathString := "/root/goProject/src/CGO/service/run2/conf.ini"
	fmt.Printf("get config path %s\n", pathString)
	LoadNewConfigFile(pathString)
	var mmap = make(map[string]string, 0)
	for _, v := range defaultConfigInstance.AllKeys() {
		mmap[v] = defaultConfigInstance.GetString(v)
	}

	v, _ := json.Marshal(mmap)
	log.Println(string(v))

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(v), &mapResult); err != nil {
		t.Fatal(err)
	}
	t.Log(mapResult)
}
