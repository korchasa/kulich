package firewall

import (
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Firewall interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*spec.OsOption) error
	FirstRun() error
	BeforeRun() error
	Add(r *spec.FirewallRule) error
	Remove(r *spec.FirewallRule) error
	AfterRun() error
}

type Config struct {
}
