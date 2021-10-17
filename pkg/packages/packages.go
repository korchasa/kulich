package packages

import (
	"github.com/korchasa/kulich/pkg/config"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Packages interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error
	FirstRun() error
	BeforeRun() error
	Add(name string) error
	Remove(name string) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
