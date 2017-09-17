package main

import (
	"./server"
    "./config"
    "./util"
    "path/filepath"
    "runtime"
	"log"
)

// description
// define default config
const (
	DefaultConfigFile  = "./default.json"
	DefaultLogFile = "proxy.log"
)

// descrition
// main 
func main() {

	// prepare to start proxy server
	log.Println("prepare to start server")
    path := util.DefaultPath()
    log.Println("homt path ---->" ,  path)
    config, err := config.Load(filepath.Join(path, DefaultConfigFile))
    if err == nil {
        runtime.GOMAXPROCS(config.MaxProcessor)
        proxy := &server.ProxyServer{}
                
        proxy.Init(config)
        proxy.Start()
    }
}
