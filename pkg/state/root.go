package state

type Root struct {
	State State `hcl:"state,block"`
}
