package services

import (
	"github.com/korchasa/ruchki/pkg/config"
	"github.com/korchasa/ruchki/pkg/sysshell"
)

type Services interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error
	FirstRun() error
	BeforeRun() error
	Add(s *Service) error
	Remove(s *Service) error
	AfterRun() error
}

type Service struct {
	Name            string
	Disabled        bool
	RestartOnChange []Watcher
}

type Watcher struct {
	Path string
	Hash string
}

type Config struct {
	DryRun bool
}
