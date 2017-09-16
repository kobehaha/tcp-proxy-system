package proxy

import (
	"../system"
	"./strategy"
	"io"
	"log"
	"net"
    "time"
)

const (
	DefaultTimeoutTime = 3
)

//description
//tcp proxy
type TcpProxy struct {
	// data *ProxyData
	strategy strategy.Strategy
	data     *ProxyData
}

//description
//set strategy
func (proxy *TcpProxy) setStrategy(name string) {
	proxy.strategy = strategy.GetStrategy(name)
	proxy.strategy.Init()
}

//description
//set backend available
func (proxy *TcpProxy) isBackendAvailable() bool {
	return true
}

//description
//dispatch
func (proxy *TcpProxy) Dispatch(con net.Conn) {
	// need check backends availabe
	if proxy.isBackendAvailable() {
		servers := "127.0.0.1:8080"
		url := proxy.strategy.Choose(con.RemoteAddr().String(), servers)
		proxy.transfer(con, url)
	} else {
		con.Close()
		log.Println("no endpoints availaber now, please check backend servers")
	}

}

func (proxy *TcpProxy) transfer(local net.Conn, remote string) {
	remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeoutTime*time.Second)
	if err != nil {
		local.Close()
		log.Println("connect to endpint error: %s ", err)
		return
	}
	sync := make(chan int, 1)
	channel := system.Channel{SrcConn:local, DstConn:remoteConn}
	go proxy.putChannel(&channel)
	go proxy.safeCopy(local, remoteConn, sync)
	go proxy.safeCopy(remoteConn, local, sync)
	go proxy.closeChannel(&channel, sync)
}

func (proxy *TcpProxy) safeCopy(from net.Conn, to net.Conn, sync chan int) {
	io.Copy(from, to)
	defer from.Close()
	sync <- 1

}

func (proxy *TcpProxy) putChannel(channel *system.Channel) {
	proxy.data.ChannelManager.Put(channel)
}

func (proxy *TcpProxy) closeChannel(channel *system.Channel, sync chan int) {
	for i := 0; i < system.ChannelPairNum; i++ {
		<-sync
	}
	proxy.data.ChannelManager.DeleteChannel(channel)

}
