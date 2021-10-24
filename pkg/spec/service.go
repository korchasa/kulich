package spec

import "fmt"

type Service struct {
	Name     string `hcl:"name,label"`
	Disabled bool   `hcl:"disabled,optional"`
}

func (s Service) Identifier() string {
	return s.Name
}

func (s Service) EqualityHash() string {
	return fmt.Sprintf("%v", s)
}

type Watcher struct {
	Path string `hcl:"path"`
	Hash string `hcl:"hash"`
}

type Services []Service

type ServicesDiff struct {
	Changed []Service
	Removed []Service
}
