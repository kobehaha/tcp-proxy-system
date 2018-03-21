package util

import (
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// description
// parse for adrress
func HostPortToAddress(host string, port uint16) string {
	return host + ":" + strconv.Itoa(int(port))
}

// description
// get slice index
func SliceIndex(slice interface{}, element interface{}) int {
	index := -1
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		return index
	}
	ev := reflect.ValueOf(element).Interface()
	length := sv.Len()
	for i := 0; i < length; i++ {
		iv := sv.Index(i).Interface()
		if reflect.DeepEqual(iv, ev) {
			index = i
			break
		}
	}
	return index

}

// description
// absolute path
func DefaultPath() string {
	absolutePath, err := filepath.Abs(".")
	if err != nil {
		log.Println("current path error:", err)
	}
	return absolutePath
}

// description
// url get to Host
func UrlToHost(url string) string {
	return strings.Split(url, ":")[0]
}

// description
// ipv4---> int
func IP4ToInt(ip string) int {
	nums := strings.Split(ip, ".")
	sum := 0
	for i := 0; i < len(nums); i++ {
		n, _ := strconv.Atoi(nums[i])
		sum += n
		sum <<= 8
	}
	return sum >> 8
}
