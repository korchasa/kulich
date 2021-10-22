package state

type Root struct {
	State   State    `hcl:"state,block"`
	Servers []Server `hcl:"server,block"`
}
