package util

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	configInstance *viper.Viper
)

func NewConfig() *viper.Viper {
	fmt.Println("init config")
	v := viper.New()
	v.AddConfigPath("/root/goProject/src/NTI-SDK/test/background/util/config.ini")
	v.ReadInConfig()
	configInstance = configInstance
	return v
}

func GetConfig() *viper.Viper {
	if configInstance == nil {
		return NewConfig()
	}
	return configInstance
}
