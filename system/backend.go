package system

import "github.com/kobehaha/tcp-proxy-system/util"

// descrition
// backend struct
type Backend struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

// description
// get backent url
func (backend *Backend) Url() string {
	return util.HostPortToAddress(backend.Host, backend.Port)
}
