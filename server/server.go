package server

import (
	"github.com/kobehaha/tcp-proxy-system/config"
	"github.com/kobehaha/tcp-proxy-system/proxy"
	"github.com/kobehaha/tcp-proxy-system/util"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	version = 5
)

// description
// define proxyServer
type ProxyServer struct {
	host             string
	port             uint16
	beattime         int
	listener         net.Listener
	requestqueuesize int
	on               bool
	proxy            proxy.Proxy
}

// description
// init
func (server *ProxyServer) Init(config *config.Config) {
	server.on = false
	server.host = config.Host
	server.port = config.Port
	server.beattime = config.Heartbeat
	server.requestqueuesize = config.RequestQueueSize
	server.setProxy(config)
}

// descritpion
// set proxy
func (server *ProxyServer) setProxy(config *config.Config) {
	// implement interface instead
	//server.proxy = &proxy.TcpProxy{}
	server.proxy = new(proxy.TcpProxy)
	server.proxy.Init(config)
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
	// start ----> heartbeat
	server.heartbeatCheckBackends()
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
			//compare channel manager count
			go server.proxy.Dispatch(con, server.requestqueuesize)

		} else {
			log.Println("client connect server error:")
		}
	}
	defer server.listener.Close()
}

// description
// stop proxy server
func (server *ProxyServer) Stop() {
	server.listener.Close()
	// server.proxy.Close()
	server.on = false
	log.Println("proxy stop success")
}

// description
// watch ---> stop signal
func (server *ProxyServer) WatchStopSignal() {
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		<-signal_channel
		server.Stop()
		log.Println("proxy recevie stop signal")
	}()
}

// description
// hearbeat to check backends server
func (server *ProxyServer) heartbeatCheckBackends() {
	timerChannel := time.NewTicker(time.Millisecond * time.Duration(server.beattime)).C
	go func() {
		for {
			select {
			case <-timerChannel:
				server.proxy.Check()
			}
		}
	}()

}
