package state

type FirewallRuleOverride struct {
	Name     string   `hcl:"name,label"`
	Ports    []string `hcl:"ports,optional"`
	Protocol *string  `hcl:"protocol,optional"`
	Targets  []string `hcl:"targets,optional"`
	IsOutput *bool    `hcl:"is_output,optional"`
}

func (o *FirewallRuleOverride) NewRule() FirewallRule {
	return FirewallRule{
		Name:     o.Name,
		Ports:    o.Ports,
		Protocol: *o.Protocol,
		Targets:  o.Targets,
		IsOutput: *o.IsOutput,
	}
}
