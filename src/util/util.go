package util

import (
	"log"
	"path/filepath"
	"reflect"
	"strconv"
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
