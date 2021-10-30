package spec

import "fmt"

type Package struct {
	Name    string
	Removed bool
}

func (p Package) Identifier() string {
	return p.Name
}

func (p Package) EqualityHash() string {
	return fmt.Sprintf("%s|%v", p.Name, p.Removed)
}

type PackagesDiff struct {
	Changed []Package
	Removed []Package
}
