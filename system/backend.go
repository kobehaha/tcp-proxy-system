package system

import "../util"

type Backend struct {
    Host string `json:"host"`
    Port unit16 `json:"port"`
}

func (backend *Backend) Url() string {
    return util.HostPortToAddress(backend.Host, backend.Port)
}
