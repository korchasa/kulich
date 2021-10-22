package state

const DefaultProtocol = "tcp"

type FirewallRule struct {
	Name       string   `hcl:"name,label"`
	Identifier string   `hcl:"identifier,optional"`
	Ports      []string `hcl:"ports"`
	Protocol   string   `hcl:"protocol,optional"`
	Targets    []string `hcl:"targets"`
	IsOutput   bool     `hcl:"is_output,optional"`
}

func (r *FirewallRule) Apply(f FirewallRuleOverride) bool {
	if r.Name != f.Name {
		return false
	}
	if f.Ports != nil {
		r.Ports = f.Ports
	}
	if f.Protocol != nil {
		r.Protocol = *f.Protocol
	}
	if f.Targets != nil {
		r.Targets = f.Targets
	}
	if f.IsOutput != nil {
		r.IsOutput = *f.IsOutput
	}
	return true
}
