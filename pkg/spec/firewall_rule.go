package spec

import "fmt"

const DefaultProtocol = "tcp"

type FirewallRule struct {
	Name     string
	Id       string
	Ports    []string
	Protocol string
	Targets  []string
	IsOutput bool
}

func (f FirewallRule) Identifier() string {
	return f.Name
}

func (f FirewallRule) EqualityHash() string {
	return fmt.Sprintf("%v", f)
}

type FirewallRulesDiff struct {
	Changed []FirewallRule
	Removed []FirewallRule
}
