package services

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Services interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*state.OsOption) error
	FirstRun() error
	BeforeRun() error
	Add(s *state.Service) error
	Remove(s *state.Service) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
