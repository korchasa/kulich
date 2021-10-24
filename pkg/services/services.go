package services

import (
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Services interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*spec.OsOption) error
	FirstRun() error
	BeforeRun() error
	Add(s *spec.Service) error
	Remove(s *spec.Service) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
