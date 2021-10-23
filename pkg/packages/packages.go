package packages

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Packages interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*state.OsOption) error
	FirstRun() error
	BeforeRun() error
	Add(p *state.Package) error
	Remove(p *state.Package) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
