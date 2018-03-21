package strategy

import (
	"time"
)

//description
//random class
type Random struct {
}

func (strategy *Random) Init() {}

//description
//random
func (strategy *Random) Choose(client string, servers []string) string {
	length := len(servers)
	url := servers[int(time.Now().UnixNano())%length]
	return url
}
