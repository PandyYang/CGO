package main

import (
	"errors"
	"fmt"
	"log"
	"ntisdk/test/cmd/v3/app"
	"ntisdk/utils"
	"runtime/cgo"
	"strings"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

// #include <stdint.h>
import "C"

// 初始化返回的句柄 包含配置信息 连接池信息 日志信息等等
type InitHandler struct {
	configInstance *viper.Viper
	redisPool      *redis.Pool
	loggerInstance *log.Logger
}

func main() {

}

//export initSDK
func initSDK(configPath *C.char) C.uintptr_t {
	config := utils.LoadConfigFile(C.GoString(configPath))
	redisPool := utils.NewRedisPool()
	logger := utils.NewLogger("")

	initHandler := InitHandler{}
	initHandler.configInstance = config
	initHandler.redisPool = redisPool
	initHandler.loggerInstance = logger
	return C.uintptr_t(cgo.NewHandle(initHandler))
}

//export destroy
func destroy(handle C.uintptr_t) {
	h := cgo.Handle(handle)
	h.Delete()
}

//export query
func query(handle C.uintptr_t, str *C.char) *C.char {
	h := cgo.Handle(handle)
	val := h.Value().(InitHandler)
	i := parseInitHandler(val)
	if i != 0 {
		fmt.Println("parse init Handler failed")
	}
	s, _ := app.Query(C.GoString(str))
	return C.CString(s)
}

func parseInitHandler(handler InitHandler) int {
	if handler.configInstance == nil || handler.redisPool == nil || handler.loggerInstance == nil {
		return -1
	}

	utils.SetConfig(handler.configInstance)
	utils.SetRedisPool(handler.redisPool)
	utils.SetLogger(handler.loggerInstance)
	return 0
}


//export downloadPackage
func downloadPackage(handle C.uintptr_t) int {
	h := cgo.Handle(handle)
	val := h.Value().(InitHandler)
	i := parseInitHandler(val)

	if i != 0 {
		return i
	}
	err := parse()
	if err != nil {
		return -1
	}
	return 0
}

//export update
func update() {
	// 升级接口 待实现
}

func parse() error {
	config := utils.GetConfig()
	logger := utils.GetLogger()

	licFilePath := config.GetString("extra.LICFILE")
	if licFilePath != "" {
		err := app.ParseCreditFile(licFilePath)
		if err != nil {
			logger.Printf("parse credit file error, %s", err.Error())
			return err
		}
		return process()
	}

	logger.Printf("check credit file failed")
	return errors.New("")
}

// //export startProcess
// func startProcess(configPath *C.char) {

// 	// 传入配置文件位置 执行程序初始化 下拉数据包等操作
// 	// path := C.GoString(configPath)
// 	config := utils.LoadConfigFile(C.GoString(configPath))

// 	if config == nil {
// 		fmt.Printf("config path can not be null")
// 		return
// 	}
// 	utils.NewLogger("")
// 	process()
// }

// //export search
// func search(str *C.char) []*C.char {
// 	rdp := utils.NewRedisPool()
// 	values := strings.Split(C.GoString(str), ",")
// 	queryResList := make([]*C.char, 0)

// 	for _, v := range values {
// 		res, _ := app.Query(v, rdp)
// 		queryResList = append(queryResList, C.CString(res))
// 	}
// 	return queryResList
// }

// 处理文件下载 解析
func process() error {

	var err error

	logger := utils.GetLogger()

	// 清理文件
	err = utils.CleanCacheDir()
	if err != nil {
		return err
	}

	// 不要在此入口开启协程 入口开启协程是针对多用户
	downloadListFileFlag := app.Download("")

	// list文件下载成功 需要进行解析操作
	if downloadListFileFlag {
		err, dataMap := app.ParseListFile()
		if err != nil {
			utils.GetLogger().Printf("parse list file failed")
			return err
		}

		if dataMap != nil && len(dataMap) > 0 {
			conf := utils.GetConfig()
			downloadScene := conf.Get("DownloadScene").([]string)
			var downloadValueList []string

			if downloadScene != nil {
				for _, v := range downloadScene {
					downloadValue := conf.Get(v)
					downloadValueList = append(downloadValueList, downloadValue.(string))
				}
			}

			// 几大威胁类型 从配置文件中读取的 加上默认的
			defaultScene := conf.Get("defaultScene").(string)
			if len(defaultScene) > 0 {
				downloadValueList = append(downloadValueList, strings.Split(defaultScene, ",")...)
			}

			if len(downloadValueList) > 0 {
				var wg sync.WaitGroup
				wg.Add(len(downloadValueList))

				// 创建可读可写的带有缓冲区的channel 用于接收多线程结果
				errChan := make(chan bool, len(downloadValueList))

				for _, v := range downloadValueList {
					var upType string
					if !strings.HasPrefix(v, "nti_") {
						upType = fmt.Sprintf("nti_%s", v)
					} else {
						upType = v
					}

					go func() {
						isDone := app.Download(upType)
						if !isDone {
							errChan <- isDone
						}
						wg.Done()
					}()
				}
				wg.Wait()

				close(errChan)
				if len(errChan) > 0 {
					logger.Printf("when handling the package, some things error, so the data not" +
						" process completely, please check the log see the details.")
					return errors.New("")
				} else {
					logger.Printf("download task finished completely.")
				}

				// handle data
				err = app.ProcessRichData()
				if err != nil {
					return err
				}

			} else {
				logger.Printf("read download pkg list failed, please check the credit xml read logic.")
				return errors.New("")
			}
		} else {
			logger.Printf("construct data map failed, please check the file LIST.xml exist in dir cache, and" +
				" have read permission.")
			return errors.New("")
		}
	}
	return errors.New("check credit key info failed")
}
