package state

type Option struct {
	Type  string `hcl:"type,label"`
	Name  string `hcl:"name,label"`
	Value string `hcl:"value"`
}
