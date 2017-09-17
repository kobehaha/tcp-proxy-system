package main

import (
	"github.com/kobehaha/tcp-proxy-system/config"
	"github.com/kobehaha/tcp-proxy-system/log"
	"github.com/kobehaha/tcp-proxy-system/server"
	"github.com/kobehaha/tcp-proxy-system/util"
	"path/filepath"
	"runtime"
)

// description
// define default config
const (
	DefaultConfigFile = "./default.json"
	DefaultLogFile    = "proxy.log"
)

// descrition
// main
func main() {

	// prepare to start proxy server
	log.Init(DefaultLogFile)
	//log.Println("prepare to start server")
	path := util.DefaultPath()
	// log.Println("homt path ---->" ,  path)
	config, err := config.Load(filepath.Join(path, DefaultConfigFile))
	if err == nil {
		runtime.GOMAXPROCS(config.MaxProcessor)
		proxy := &server.ProxyServer{}
		proxy.Init(config)
		proxy.WatchStopSignal()
		proxy.Start()
	}
}
