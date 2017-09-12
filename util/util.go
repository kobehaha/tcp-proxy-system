package util

import (
	"strconv"
)

// descrition
// parse for adrress
func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
