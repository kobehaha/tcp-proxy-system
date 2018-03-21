package proxy

import (
	"net"
	"github.com/kobehaha/tcp-proxy-system/config"
)

//description
//proxy interface
type Proxy interface {
	Init(config *config.Config)
	Dispatch(con net.Conn, request_size int)
	Check()
}
