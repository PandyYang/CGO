package main

import "C"

import (
	"fmt"
	"log"
	"net/http"
	"ntisdk/test/background/util"

	"github.com/kardianos/service"
)

var logger service.Logger

// Program structures.
//
//	Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	util.NewConfig()
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() error {
	logger.Infof("I'm running %v.", service.Platform())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

//export WaitProcess
func WaitProcess(name string) {
	config := util.GetConfig()
	res := config.GetString("test.123")
	fmt.Println(name + res)
}

func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

var ss service.Service

//export Cmd
func Cmd(flag string) {

	if len(flag) != 0 && ss != nil {
		err := service.Control(ss, flag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}

		return
	}
}

func main() {
	// svcFlag := flag.String("service", "", "Control the system service.")
	// flag.Parse()

	// options := make(service.KeyValue)
	// options["Restart"] = "on-success"
	// options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	// svcConfig := &service.Config{
	// 	Name:         "GoExample",
	// 	DisplayName:  "Go Example",
	// 	Description:  "This is a Go application.",
	// 	Dependencies: []string{},
	// 	// "Requires=network.target",
	// 	// "After=network-online.target syslog.target"
	// 	Option: options,
	// }

	// prg := &program{}
	// s, err := service.New(prg, svcConfig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// errs := make(chan error, 5)
	// logger, err = s.Logger(errs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// go func() {
	// 	for {
	// 		err := <-errs
	// 		if err != nil {
	// 			log.Print(err)
	// 		}
	// 	}
	// }()

	// if len(*svcFlag) != 0 {
	// 	ss = s
	// 	err := service.Control(s, *svcFlag)
	// 	if err != nil {
	// 		log.Printf("Valid actions: %q\n", service.ControlAction)
	// 		log.Fatal(err)
	// 	}

	// 	return
	// }

	// err = s.Run()
	// if err != nil {
	// 	logger.Error(err)
	// }

}
