package state

type State struct {
	Role Role `hcl:"role,block"`
}
