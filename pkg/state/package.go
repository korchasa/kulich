package state

import "fmt"

type Package struct {
	Name    string `hcl:"name,label"`
	Removed bool   `hcl:"removed,optional"`
}

func (p Package) Identifier() string {
	return p.Name
}

func (p Package) EqualityHash() string {
	return fmt.Sprintf("%s|%v", p.Name, p.Removed)
}

type Packages []Package

type PackagesDiff struct {
	Changed []Package
	Removed []Package
}
