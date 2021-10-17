package os

import (
	"github.com/korchasa/kulich/pkg/config"
	"github.com/korchasa/kulich/pkg/filesystem"
	"github.com/korchasa/kulich/pkg/firewall"
	"github.com/korchasa/kulich/pkg/packages"
	"github.com/korchasa/kulich/pkg/services"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Os interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error
	FirstRun() error
	BeforeAll() error
	AddUser(u *User) error
	RemoveUser(u *User) error
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
	Name   string
	System bool
}
