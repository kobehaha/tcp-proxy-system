package proxy

import (
	"net"
)

//description
//proxy interface
type Proxy interface {
	Dispatch(con net.Conn)
}
