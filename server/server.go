package server

import (
	"../util"
	"log"
	"net"
)

type ProxyServer struct {
	host     string
	port     uint16
	beatime  int
	listener net.Listener
	on       bool
}

func (server *ProxyServer) Init() {
	server.on = false
	server.host = "127.0.0.1"
	server.port = 1000
}

func (server *ProxyServer) Address() string {
	return util.HostPortToAddress(server.host, server.port)
}

func (server *ProxyServer) Start() {
	log.Println("start server and server address is " + server.Address())
	local, err := net.Listen("tcp", server.Address())
	if err != nil {
		log.Panic("server start error" + err.Error())
	}
	log.Println("proxy server start ok")
	server.listener = local
	server.on = true
	for server.on {
		_, err := server.listener.Accept()
		if err == nil {
			log.Println("start -----> dispatch")
		} else {
			log.Println("client connect server error:")
		}
	}
	defer server.listener.Close()
}

func (server *ProxyServer) Stop() {
	server.listener.Close()
	// server.proxy.Close()
	server.on = false
	log.Println("proxy stop success")
}
