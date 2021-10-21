package state

type UserOverride struct {
	Name   string `hcl:"name,label"`
	System *bool  `hcl:"system,optional"`
}

func (o *UserOverride) NewUser() User {
	return User{
		Name:   o.Name,
		System: *o.System,
	}
}
