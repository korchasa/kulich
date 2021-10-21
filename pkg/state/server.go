package state

type Server struct {
	Name         string                `hcl:"name,label"`
	System       *SystemOverride       `hcl:"system,block"`
	Applications []ApplicationOverride `hcl:"application,block"`
}
