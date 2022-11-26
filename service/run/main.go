package main

import (
	"context"
	"fmt"
	"github.com/kardianos/service"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

type services struct {
	log service.Logger
	srv *http.Server
	cfg *service.Config
}

func (srv *services) Start(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start run http server")
	}

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		return err
	}
	go srv.srv.Serve(lis)
	loadConf()
	log.Println("load conf done.")
	return nil
}

func (srv *services) Stop(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start stop http server")
	}
	return srv.srv.Shutdown(context.Background())
}

func readFile(str string) {
	log.Printf("read file %s", str)
	conf := Config()
	res := conf.GetString("test.123")
	log.Printf("load conf %s", res)
}

func loadConf() {
	LoadNewConfigFile("E:\\GolandProject\\CGO\\service\\run\\conf.ini")
}

func main() {

	fFile := pflag.StringP("file", "F", "", "file to read, just test to load the certificate file")
	pflag.Parse()

	if len(*fFile) > 0 {
		readFile(*fFile)
		return
	}

	File, err := os.Create("http-server.log")
	if err != nil {
		File = os.Stdout
	}
	defer File.Close()

	log.SetOutput(File)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path)
	})

	var s = &services{srv: &http.Server{Handler: http.DefaultServeMux}, cfg: &service.Config{
		Name:        "GoHttpServer",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}}

	sys := service.ChosenSystem()
	srv, err := sys.New(s, s.cfg)
	if err != nil {
		log.Fatalf("Init service error:%s\n", err.Error())
	}

	s.log, err = srv.SystemLogger(nil)
	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := srv.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
		case "uninstall":
			err := srv.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
		}
		return
	}

	err = srv.Run()
	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}

func command() {

}

const (
	defaultConfigName = "nti-sdk-config"
	ConfigPath        = "/root/goProject/src/NTI-SDK/conf"
)

var (
	defaultConfigInstance *viper.Viper
	onceConfig            sync.Once

	DefaultPkgDownloadValue = map[string]interface{}{
		"ip":     "watchlist_ip",
		"domain": "watchlist_domain",
		"url":    "watchlist_url",
		"sample": "watchlist_sample",
	}

	// DefaultPkgDownloadScene "mining,apt,blackmail,highPrecision"
	DefaultPkgDownloadScene = map[string]interface{}{
		"mining":        "nti_miner_ioc",
		"apt":           "nti_threat_actor",
		"blackmail":     "nti_ransomware",
		"highPrecision": "nti_b_200m",
	}
)

type MyViper struct {
	viper.Viper
}

// NewConfig : Create new config provider
func NewConfig(name string) *viper.Viper {
	if defaultConfigInstance != nil {
		return defaultConfigInstance
	}

	fmt.Println("init new config")

	if name == "" {
		name = defaultConfigName
	}

	defer afterInit()

	c := viper.New()
	c.AutomaticEnv()
	c.SetConfigType("ini")
	c.AddConfigPath(ConfigPath)
	c.Set("123", "456")

	err := c.ReadInConfig()
	if err != nil {
		return nil
	}

	// once可以防止内存可见性问题
	onceConfig.Do(func() {
		if defaultConfigInstance == nil {
			defaultConfigInstance = c
		}
	})

	return defaultConfigInstance
}

func Config() *viper.Viper {
	if defaultConfigInstance != nil {
		return defaultConfigInstance
	}
	log.Println("config is nil")
	return nil
}

func SetDefaultConfig(defaults map[string]interface{}) error {
	if defaultConfigInstance == nil || defaults == nil {
		return fmt.Errorf("no viper instance, NewConfig first")
	}

	for k, v := range defaults {
		if v != nil {
			defaultConfigInstance.SetDefault(k, v)
		}
	}

	return nil
}

func LoadNewConfigFile(path string) *viper.Viper {

	var v = viper.New()
	v.SetConfigFile(path)
	v.ReadInConfig()
	defaultConfigInstance = v
	return defaultConfigInstance

	//if defaultConfigInstance == nil {
	//	defaultConfigInstance = NewConfig("")
	//}
	//onceConfig.Do(func() {
	//	defaultConfigInstance.SetConfigFile(path)
	//	defaultConfigInstance.MergeInConfig()
	//})
}

func afterInit() {

	if defaultConfigInstance == nil {
		NewConfig("")
	} else {
		err := SetDefaultConfig(DefaultPkgDownloadValue)
		if err != nil {
			return
		}

		err = SetDefaultConfig(DefaultPkgDownloadScene)
		if err != nil {
		}

		// 从配置文件中读取默认的类型
		defaultScene := defaultConfigInstance.Get("nti.DEFAULT_DOWNLOAD_SCENE").(string)
		if len(defaultScene) > 0 {
			defaultConfigInstance.Set("defaultScene", defaultScene)
		}
	}
}
