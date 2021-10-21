package state

type User struct {
	Name   string `hcl:"name,label"`
	System bool   `hcl:"system,optional"`
}

func (u *User) Apply(uo UserOverride) bool {
	if u.Name != uo.Name {
		return false
	}
	if uo.System != nil {
		u.System = *uo.System
	}
	return true
}
