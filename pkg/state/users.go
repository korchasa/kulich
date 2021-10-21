package state

type Users []User

func (u *Users) Apply(override []UserOverride) {
	for _, oo := range override {
		applied := false
		for i := range *u {
			applied = applied || (*u)[i].Apply(oo)
		}
		if !applied {
			*u = append(*u, oo.NewUser())
		}
	}
}
