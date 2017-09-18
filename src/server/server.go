package server
import (
	"../config"
	"../proxy"
	"../util"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = 5
)

// description
// define proxyServer
type ProxyServer struct {
	host     string
	port     uint16
	beattime int
	listener net.Listener
    requestqueuesize int 
	on       bool
	proxy    *proxy.TcpProxy
}

// description
// init
func (server *ProxyServer) Init(config *config.Config) {
	server.on = false
	//server.host = "127.0.0.1"
	//server.port = 1000
	server.host = config.Host
	server.port = config.Port
	server.beattime = config.Hearbeat
    server.requestqueuesize = config.RequestQueueSize
	server.setProxy(config)
}

// descritpion
// set proxy
func (server *ProxyServer) setProxy(config *config.Config) {
	server.proxy = &proxy.TcpProxy{}
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

// descrition
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
