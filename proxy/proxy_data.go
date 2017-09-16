package proxy

import (
	"../system"
	"sync"
)

type ProxyData struct {
	Service        string
	Host           string
	Port           uint16
	Backends       map[string]system.Backend
	Deads          map[string]system.Backend
	ChannelManager *system.ChannelManager
	mutex          *sync.RWMutex
}
