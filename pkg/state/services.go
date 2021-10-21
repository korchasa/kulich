package state

type Services []Service

func (s *Services) Apply(override []ServiceOverride) {
	for _, oo := range override {
		applied := false
		for i := range *s {
			applied = applied || (*s)[i].Apply(oo)
		}
		if !applied {
			*s = append(*s, oo.NewService())
		}
	}
}
