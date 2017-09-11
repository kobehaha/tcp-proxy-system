package main

import (
	"./server"
	"log"
)

const (
	DefaultConfig  = "config/default.json"
	DefaultLogFile = "proxy.log"
)

func main() {

	// prepare to start proxy server
	log.Println("prepare to start server")
	proxy := &server.ProxyServer{}
	proxy.Init()
	proxy.Start()

}
