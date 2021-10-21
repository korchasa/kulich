package state

type Package struct {
	Name    string `hcl:"name,label"`
	Removed bool   `hcl:"removed,optional"`
}

func (p *Package) Apply(po PackageOverride) bool {
	if p.Name != po.Name {
		return false
	}
	if po.Removed != nil {
		p.Removed = *po.Removed
	}
	return true
}

type PackageOverride struct {
	Name    string `hcl:"name,label"`
	Removed *bool  `hcl:"removed"`
}

func (o *PackageOverride) NewPackage() Package {
	return Package{
		Name:    o.Name,
		Removed: *o.Removed,
	}
}
