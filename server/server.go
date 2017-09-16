package server

import (
	"../util"
    "../proxy"
	"log"
	"net"
)

// description
// define proxyServer 
type ProxyServer struct {
	host     string
	port     uint16
	beatime  int
	listener net.Listener
	on       bool
    proxy    proxy.Proxy
}

// description
// init 
func (server *ProxyServer) Init() {
	server.on = false
	server.host = "127.0.0.1"
	server.port = 1000
}

// description
// read and parse from config file
func (server *ProxyServer) LoadConfg() bool {
    return true
}

// description
// get address from proxyServer 
func (server *ProxyServer) Address() string {
	return util.HostPortToAddress(server.host, server.port)
}

// description
// start proxy
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
		con, err := server.listener.Accept()
		if err == nil {
			log.Println("start -----> dispatch")
            go server.proxy.Dispatch(con)

		} else {
			log.Println("client connect server error:")
		}
	}
	defer server.listener.Close()
}


// descrition
// stop proxy server
func (server *ProxyServer) Stop() {
	server.listener.Close()
	// server.proxy.Close()
	server.on = false
	log.Println("proxy stop success")
}
