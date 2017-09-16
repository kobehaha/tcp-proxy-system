package proxy

import (
	"../system"
	"sync"
)

// description
// proxyData 
type ProxyData struct {
	Service        string
	Host           string
	Port           uint16
	Backends       map[string]system.Backend
	Deads          map[string]system.Backend
	ChannelManager *system.ChannelManager
	mutex          *sync.RWMutex
}


// description
// proxyData init from config file
func (proxyData *ProxyData) Init(){
   proxyData.ChannelManager = new(system.ChannelManager) 
   proxyData.ChannelManager.Init()
   proxyData.mutex = new(sync.RWMutex)
}
