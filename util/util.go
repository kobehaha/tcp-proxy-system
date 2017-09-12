package util

import (
	"strconv"
)

// description
// parse for adrress
func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}
