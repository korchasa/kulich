package state

type Rule struct {
	Name       string   `hcl:"name,label"`
	Identifier string   `hcl:"identifier,optional"`
	Ports      []string `hcl:"ports"`
	Protocol   string   `hcl:"protocol,optional"`
	Targets    []string `hcl:"targets"`
	IsOutput   bool     `hcl:"is_output,optional"`
}

const DefaultProtocol = "tcp"

type RuleOverride struct {
	Name       string    `hcl:"name,label"`
	Identifier *string   `hcl:"identifier"`
	Ports      *[]string `hcl:"ports"`
	Protocol   *string   `hcl:"protocol"`
	Targets    *[]string `hcl:"targets"`
	IsOutput   *bool     `hcl:"is_output"`
}
