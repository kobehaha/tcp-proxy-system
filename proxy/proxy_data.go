package proxy

import (
	"../system"
    "../config"
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
func (proxyData *ProxyData) Init(config *config.Config){
   proxyData.Service = config.Service
   proxyData.Host = config.Host
   proxyData.Port = config.Port
   proxyData.ChannelManager = new(system.ChannelManager) 
   proxyData.ChannelManager.Init()
   proxyData.setBackends(config.Backends)
   proxyData.mutex = new(sync.RWMutex)
}

// description
// set backends
func (proxyData *ProxyData) setBackends(backends []system.Backend){
    proxyData.Backends = make(map[string]system.Backend)
    for _, backend := range backends{
        proxyData.Backends[backend.Url()] = backend
    }
}


// description
// set backendUrl
func (proxyData *ProxyData) BackendUrls()  []string {
    proxyData.mutex.RLock()
    defer proxyData.mutex.RUnlock()
    _map := proxyData.Backends
    keys := make([]string, 0, len(_map))
    for k := range _map {
        keys = append(keys, k)
    }
    return keys
}
