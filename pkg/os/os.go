package os

import (
	"github.com/korchasa/ruchki/pkg/config"
	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/services"
	"github.com/korchasa/ruchki/pkg/sysshell"
)

type Os interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error
	FirstRun() error
	BeforeAll() error
	AddUser(u *User)
	RemoveUser(u *User)
	SetOption(opt *config.Option) error
	BeforePackages(p *packages.Packages) error
	AfterPackages(p *packages.Packages) error
	BeforeFilesystem(f *filesystem.Filesystem) error
	AfterFilesystem(f *filesystem.Filesystem) error
	BeforeServices(f *services.Services) error
	AfterServices(f *services.Services) error
	BeforeFirewall(f *firewall.Firewall) error
	AfterFirewall(f *firewall.Firewall) error
	AfterAll() error
}

type User struct {
	Name           string
	Shell          string
	Home           string
	AuthorizedKeys string
	Removed        bool
}
