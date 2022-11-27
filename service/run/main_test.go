package main

import "testing"

func TestReadConfig(t *testing.T) {
	v := LoadNewConfigFile("/root/goProject/src/CGO/service/run/conf.ini")
	vv := v.GetString("123.key")
	println(vv)
}
