package state

import "fmt"

const DefaultProtocol = "tcp"

type FirewallRule struct {
	Name     string   `hcl:"name,label"`
	Id       string   `hcl:"identifier,optional"`
	Ports    []string `hcl:"ports"`
	Protocol string   `hcl:"protocol,optional"`
	Targets  []string `hcl:"targets"`
	IsOutput bool     `hcl:"is_output,optional"`
}

func (f FirewallRule) Identifier() string {
	return f.Name
}

func (f FirewallRule) EqualityHash() string {
	return fmt.Sprintf("%v", f)
}

type FirewallRules []FirewallRule

type FirewallRulesDiff struct {
	Changed []FirewallRule
	Removed []FirewallRule
}
