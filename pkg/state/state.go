package state

type State struct {
	Name         string       `hcl:"name,label"`
	Config       Config       `hcl:"config,block"`
	System       System       `hcl:"system,block"`
	Applications Applications `hcl:"application,block"`
}

func (r *State) Apply(override Server) bool {
	r.System.Apply(override.System)
	r.Applications.Apply(override.Applications)
	return true
}
