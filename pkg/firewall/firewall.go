package firewall

type Firewall interface {
	Setup(c *Config) error
	Add(r *Rule) error
	Remove(r *Rule) error
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
