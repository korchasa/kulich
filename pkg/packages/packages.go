package packages

import (
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Packages interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*spec.OsOption) error
	FirstRun() error
	BeforeRun() error
	Add(p *spec.Package) error
	Remove(p *spec.Package) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
