package strategy

const (
	PollName   = "poll"
	IpHashName = "iphash"
	RandomName = "random"
)

var register = make(map[string]Strategy)

//descriptioin
//strategy
type Strategy interface {
	Init()
	Choose(client string, server []string) string
}

//description
//init
func init() {
	register[RandomName] = new(Random)
	register[IpHashName] = new(IpHash)
}

//description
//get strategy
func GetStrategy(name string) Strategy {
	return register[name]
}
