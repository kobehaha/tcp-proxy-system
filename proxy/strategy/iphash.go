package strategy

import (
	"github.com/kobehaha/tcp-proxy-system/util"
)

type IpHash struct {
}

// description
// Init
func (strategy *IpHash) Init() {}


// description
// Choose
func (strategy *IpHash) Choose(client string, servers []string)  string {
	ip := util.UrlToHost(client)
	intIp := util.IP4ToInt(ip)
	length := len(servers)
	url := servers[intIp%length]
	return url
}
