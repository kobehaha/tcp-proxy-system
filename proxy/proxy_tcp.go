package proxy

import (
	"../system"
	"./strategy"
	//	"io"
	"log"
	"net"
	"time"
)

const (
	DefaultTimeoutTime = 3
)

// description
// tcp proxy
type TcpProxy struct {
	// data *ProxyData
	strategy strategy.Strategy
	data     *ProxyData
}

// description
// init TcpProxy
func (proxy *TcpProxy) Init() {
	proxy.setProxyData(&ProxyData{})
	proxy.setStrategy("random")
}

// description
// set strategy
func (proxy *TcpProxy) setStrategy(name string) {
	proxy.strategy = strategy.GetStrategy(name)
	proxy.strategy.Init()
}

// description
// tcp proxy set data
func (proxy *TcpProxy) setProxyData(proxyData *ProxyData) {
	proxy.data = proxyData
	proxy.data.Init()
}

// description
// set backend available
func (proxy *TcpProxy) isBackendAvailable() bool {
	return true
}

// description
// dispatch
func (proxy *TcpProxy) Dispatch(con net.Conn) {
	// need check backends availabe
	log.Println("check availabe backends")
	if proxy.isBackendAvailable() {
        //servers := []string{"192.168.33.19:8000"}
        servers := []string{"127.0.0.1:21288"}
		// set static ---> change dynamic
		url := proxy.strategy.Choose(con.RemoteAddr().String(), servers)
		proxy.transfer(con, url)
	} else {
		con.Close()
		log.Println("no endpoints availaber now, please check backend servers")
	}

}

// description
// transfer ---> client---> proxy--->backends
func (proxy *TcpProxy) transfer(local net.Conn, remote string) {
	remoteConn, err := net.DialTimeout("tcp", remote, DefaultTimeoutTime*time.Second)
	if err != nil {
		local.Close()
		log.Println("connect to endpint error: %s ", err)
		return
	}
	sync := make(chan int, 1)
	channel := system.Channel{SrcConnection: local, DstConnection: remoteConn}
	// set client 3 minite for upating , force close
	local.SetReadDeadline(time.Now().Add(time.Minute * 3))
	local.SetWriteDeadline(time.Now().Add(time.Minute * 3))
	// set proxy ---> for byte
	go proxy.putChannel(&channel)
	go proxy.safeCopy(local, remoteConn, sync)
	go proxy.safeCopy(remoteConn, local, sync)
	go proxy.closeChannel(&channel, sync)
}

// description
// safe Copy --->[]byte
// method 1 ----> parse byte from conn ---| for more designer, I choose method 1 
// method 2 ----> just use io.Copy(from, to)
func (proxy *TcpProxy) safeCopy(from net.Conn, to net.Conn, sync chan int) {
    // method 1
	for {
		buf := make([]byte, 512)
		n, err := from.Read(buf)
		if err != nil {
			log.Println("read error---------->:", err)
			break
		}
		if n == 0 {
			sync <- 1
			break
		}
		log.Println("read:", string(buf[:n]))
		// check --> byte
		to.Write(buf)
		log.Println("write:--->")
	}

    // method 2 
	//   io.Copy(from, to)
	// check --> byte

	defer from.Close()
	//sync <- 1

}

// description
// put Channel ---> manager
func (proxy *TcpProxy) putChannel(channel *system.Channel) {
	proxy.data.ChannelManager.Put(channel)
}

// description
// close Channel
func (proxy *TcpProxy) closeChannel(channel *system.Channel, sync chan int) {
	for i := 0; i < system.ChannelPairNum; i++ {
		<-sync
	}
	proxy.data.ChannelManager.Delete(channel)

}
