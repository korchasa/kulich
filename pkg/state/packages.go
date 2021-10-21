package state

type Packages []Package

func (p *Packages) Apply(override []PackageOverride) {
	for _, oo := range override {
		applied := false
		for i := range *p {
			applied = applied || (*p)[i].Apply(oo)
		}
		if !applied {
			*p = append(*p, oo.NewPackage())
		}
	}
}
