package main

import (
	"./server"
	"log"
)

// description
// define default config
const (
	DefaultConfig  = "config/default.json"
	DefaultLogFile = "proxy.log"
)

// descrition
// main 
func main() {

	// prepare to start proxy server
	log.Println("prepare to start server")
    proxy := &server.ProxyServer{}
	proxy.Init()
	proxy.Start()

}
