package proxy

import (
	"../config"
	"../system"
	"./strategy"
	"log"
	"net"
	"time"
)

const (
	DefaultTimeoutTime = 3
)

// description/
// tcp proxy
type TcpProxy struct {
	// data *ProxyData
	strategy strategy.Strategy
	data     *ProxyData
}

// description
// init TcpProxy
func (proxy *TcpProxy) Init(config *config.Config) {
	//proxy.setProxyData(&ProxyData{})
	proxy.data = new(ProxyData)
	proxy.data.Init(config)
	proxy.setStrategy(config.Strategy)
}

// description
// set strategy
func (proxy *TcpProxy) setStrategy(name string) {
	proxy.strategy = strategy.GetStrategy(name)
	proxy.strategy.Init()
}

// description
// set backend available
func (proxy *TcpProxy) isBackendAvailable() bool {
	return len(proxy.data.Backends) > 0
}

// description
// dispatch
func (proxy *TcpProxy) Dispatch(con net.Conn, requestqueuesize int) {

	log.Println("current request sie ----> ", proxy.data.getRequestSrcLen)
	// compare channalManager count
	if proxy.data.getRequestSrcLen() >= requestqueuesize {
		// need to add ---> requesting queue | channel notify
		log.Println("new request need to wait for requestqueue")
		con.Close()
		return
	}
	// need check backends availabe
	log.Println("check availabe backends")
	if proxy.isBackendAvailable() {
		//servers := []string{"192.168.33.19:8000"}
		//servers := []string{"127.0.0.1:21288"}
		servers := proxy.data.BackendUrls()
		log.Println("servers----> ", servers)
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
	log.Println("remote --->", remote)
	if err != nil {
		local.Close()
		log.Println("connect to endpint error:  ", err)
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

// description
// check backends
func (proxy *TcpProxy) Check() {
	for _, backend := range proxy.data.Backends {
		_, err := net.Dial("tcp", backend.Url())
		if err != nil {
			proxy.data.Clear(backend.Url())
			log.Println("clean backendUrl which is not available -----> ", backend.Url())
		}
	}

	for _, backend := range proxy.data.Deads {
		_, err := net.Dial("tcp", backend.Url())
		if err == nil {
			proxy.data.Recover(backend.Url())
			log.Println("recover backendUrl which is availabe ----->", backend.Url())
		}
	}

}
