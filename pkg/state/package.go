package state

type Package struct {
	Name    string `hcl:"name,label"`
	Removed bool   `hcl:"removed,optional"`
}
