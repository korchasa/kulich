package firewall

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Firewall interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*state.Option) error
	FirstRun() error
	BeforeRun() error
	Add(r *state.Rule) error
	Remove(r *state.Rule) error
	AfterRun() error
}

type Config struct {
}
