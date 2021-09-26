package firewall

type FirewallRule struct {
	Port int
	Sources []string
	Output bool
}

type FirewallDriverConfig struct {
	Driver string
	EnabledInput bool
	EnabledOutput bool
}

type FirewallDriver interface {
	Setup(c *FirewallDriverConfig) error
	ApplyRule(f *FirewallRule) error
}
