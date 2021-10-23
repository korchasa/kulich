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
