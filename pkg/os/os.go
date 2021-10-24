package os

import (
	"github.com/korchasa/kulich/pkg/filesystem"
	"github.com/korchasa/kulich/pkg/firewall"
	"github.com/korchasa/kulich/pkg/packages"
	"github.com/korchasa/kulich/pkg/services"
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/korchasa/kulich/pkg/sysshell"
)

type Os interface {
	Config(dryRun bool, sh sysshell.Sysshell, opts ...*spec.OsOption) error
	FirstRun() error
	BeforeAll() error
	AddUser(u *spec.User) error
	RemoveUser(u *spec.User) error
	SetOption(opt *spec.OsOption) error
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
