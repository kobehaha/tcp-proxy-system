package proxy

import (
	"../config"
	"net"
)

//description
//proxy interface
type Proxy interface {
	Init(config *config.Config)
	Dispatch(con net.Conn, request_size int)
	Check()
}
