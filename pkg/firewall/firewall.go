package firewall

type Firewall interface {
	FirstRun() error
	BeforeRun() error
	Add(r *Rule) error
	Remove(r *Rule) error
	AfterRun() error
}

type Rule struct {
	Identifier string
	Ports      []string
	Protocol   string
	Targets    []string
	IsOutput   bool
}

const DefaultProtocol = "tcp"

type Config struct {
}
