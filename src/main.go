package main

import (
	"./config"
	"./log"
	"./server"
	"./util"
	"fmt"
	"path/filepath"
	"runtime"
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

	//log.Println("prepare to start server")
	path := util.DefaultPath()
	// log.Println("homt path ---->" ,  path)
	config, err := config.Load(filepath.Join(path, DefaultConfigFile))
	init_log(config)

	if err == nil {
		runtime.GOMAXPROCS(config.MaxProcessor)
		proxy := &server.ProxyServer{}
		proxy.Init(config)
		proxy.WatchStopSignal()
		proxy.Start()
	}
}

func init_log(config *config.Config) {
	// log
	logfile := config.LogFile
	fmt.Println("logfile = ", logfile)
	if logfile == "" {
		log.Init(DefaultLogName, DefaultLogFileLocation)
	} else {
		log.Init(DefaultLogName, logfile)
	}

}
