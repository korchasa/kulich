package firewall

import (
	"github.com/korchasa/ruchki/pkg/config"
	"github.com/korchasa/ruchki/pkg/sysshell"
)

type Firewall interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error
	FirstRun() error
	BeforeRun() error
	Add(r *Rule) error
	Remove(r *Rule) error
	AfterRun() error
}

type Rule struct {
	Identifier string
	Ports      []string
	Protocol   string
	Targets    []string
	IsOutput   bool
}

const DefaultProtocol = "tcp"

type Config struct {
}
