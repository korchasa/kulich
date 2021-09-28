package firewall

type Driver interface {
	Setup(c *DriverConfig) error
	ApplyRule(f *Rule) error
}

type Rule struct {
	Port int
	Sources []string
	Output bool
}

type DriverConfig struct {
	Driver string
	EnabledInput bool
	EnabledOutput bool
}
