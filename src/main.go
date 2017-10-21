package main

import (
	"./config"
	"./log"
	"./server"
	"./util"
	"flag"
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
)

// descrition
// main
func main() {

	var cmdConfigFile string
	var configInfo *config.Config
	var err interface{}
	var defaultPath string

	flag.StringVar(&cmdConfigFile, "conf", DefaultConfigFile, "--conf config path location")
	flag.Parse()

	if strings.Compare(cmdConfigFile, DefaultConfigFile) == 0 {
		defaultPath = util.DefaultPath()
		configInfo, err = config.Load(filepath.Join(defaultPath, DefaultConfigFile))
	} else {
		configInfo, err = config.Load(cmdConfigFile)
	}
	init_log(configInfo)

	if err == nil {
		runtime.GOMAXPROCS(configInfo.MaxProcessor)
		proxy := &server.ProxyServer{}
		proxy.Init(configInfo)
		proxy.WatchStopSignal()
		proxy.Start()
	}
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
