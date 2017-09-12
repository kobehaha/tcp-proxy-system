package proxy

import (
    "net"
)

const (
    DefaultTimeoutTime = 3
)
 
//description
//tcp proxy
type TcpProxy struct {
    data *ProxyData
}
