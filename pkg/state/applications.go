package state

type Applications []Application

func (a *Applications) Apply(override []ApplicationOverride) {
	for _, oo := range override {
		applied := false
		for i := range *a {
			applied = applied || (*a)[i].Apply(oo)
		}
		if !applied {
			*a = append(*a, oo.NewApplication())
		}
	}
}
