package state

type Role struct {
	Name         string       `hcl:"name,label"`
	Config       Config       `hcl:"config,block"`
	System       System       `hcl:"system,block"`
	Applications Applications `hcl:"application,block"`
	Servers      []Server     `hcl:"server,block"`
}

func (r *Role) Apply(override Server) bool {
	r.System.Apply(override.System)
	r.Applications.Apply(override.Applications)
	r.Servers = nil
	return true
}
