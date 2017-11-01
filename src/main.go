package main

import (
	"./config"
	"./hack"
	"./log"
	"./server"
	"./util"
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// description
// define default config
const (
	DefaultConfigFile      = "../conf/default.json"
	DefaultLogFileLocation = "../logs"
	DefaultLogName         = "proxy.log"
	Start                  = "start"
	Stop                   = "stop"
	Status                 = "status"
)

const banner string = `
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx WELCOME USE TCP PROXY xxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
`

// descrition
// main
func main() {

	var cmdConfigFile string
	var cmd string
	var configInfo *config.Config
	var err interface{}
	var defaultPath string

	flag.StringVar(&cmdConfigFile, "conf", DefaultConfigFile, "-conf config path location")
	flag.StringVar(&cmd, "s", "start", "-s define start|stop|status")
	flag.Parse()

	fmt.Println(banner)
	fmt.Printf("version : %s\n", hack.Version)
	fmt.Printf("Build Info: %s\n", hack.Compile)

	if strings.Compare(cmdConfigFile, DefaultConfigFile) == 0 {
		defaultPath = util.DefaultPath()
		configInfo, err = config.Load(filepath.Join(defaultPath, DefaultConfigFile))
	} else {
		configInfo, err = config.Load(cmdConfigFile)
	}
	if err != nil {
		fmt.Println("tcp proxy boostrap failed : %s\n", err)
	}

	init_log(configInfo)
	runtime.GOMAXPROCS(configInfo.MaxProcessor)
	proxy := &server.ProxyServer{}
	proxy.Init(configInfo)
	proxy.WatchStopSignal()
	proxy.Start()

}

func init_log(config *config.Config) {
	// log
	logfile := config.LogFile
	if logfile == "" {
		log.Init(DefaultLogName, DefaultLogFileLocation)
	} else {
		log.Init(DefaultLogName, logfile)
	}

}
