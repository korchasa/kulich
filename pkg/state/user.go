package state

type User struct {
	Name   string `hcl:"name,label"`
	System bool   `hcl:"system,optional"`
}
