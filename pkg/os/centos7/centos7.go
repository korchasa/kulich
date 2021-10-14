package centos7

import (
	"fmt"
	"github.com/korchasa/ruchki/pkg/config"
	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/os"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/services"
	"github.com/korchasa/ruchki/pkg/sysshell"
)

type Centos7 struct {
	dryRun bool
	sh     sysshell.Sysshell
}

func (c *Centos7) Config(dryRun bool, sh sysshell.Sysshell, opts ...*config.Option) error {
	c.sh = sh
	c.dryRun = dryRun
	for _, v := range opts {
		switch v.Type {
		default:
			return fmt.Errorf("unsupported option type `%s`", v.Type)
		}
	}

	return nil
}

func (c Centos7) FirstRun() error {
	panic("implement me")
}

func (c Centos7) BeforeAll() error {
	panic("implement me")
}

func (c Centos7) AddUser(u *os.User) {
	panic("implement me")
}

func (c Centos7) RemoveUser(u *os.User) {
	panic("implement me")
}

func (c Centos7) SetOption(opt *config.Option) error {
	panic("implement me")
}

func (c Centos7) BeforePackages(p *packages.Packages) error {
	panic("implement me")
}

func (c Centos7) AfterPackages(p *packages.Packages) error {
	panic("implement me")
}

func (c Centos7) BeforeFilesystem(f *filesystem.Filesystem) error {
	panic("implement me")
}

func (c Centos7) AfterFilesystem(f *filesystem.Filesystem) error {
	panic("implement me")
}

func (c Centos7) BeforeServices(f *services.Services) error {
	panic("implement me")
}

func (c Centos7) AfterServices(f *services.Services) error {
	panic("implement me")
}

func (c Centos7) BeforeFirewall(f *firewall.Firewall) error {
	panic("implement me")
}

func (c Centos7) AfterFirewall(f *firewall.Firewall) error {
	panic("implement me")
}

func (c Centos7) AfterAll() error {
	panic("implement me")
}
