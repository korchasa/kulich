package pkg

import (
	"fmt"
	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/korchasa/ruchki/pkg/filesystem/posix"
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/firewall/iptables"
	"github.com/korchasa/ruchki/pkg/os"
	"github.com/korchasa/ruchki/pkg/os/centos7"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/packages/yum"
	"github.com/korchasa/ruchki/pkg/services"
	"github.com/korchasa/ruchki/pkg/services/systemd"
)

func NewFilesystem(name string) (filesystem.Filesystem, error) {
	switch name {
	case "posix":
		return new(posix.Posix), nil
	}
	return nil, fmt.Errorf("unsupported filesystem type `%s`", name)
}

func NewFirewall(name string) (firewall.Firewall, error) {
	switch name {
	case "iptables":
		return new(iptables.Iptables), nil
	}
	return nil, fmt.Errorf("unsupported firewall `%s`", name)
}

func NewOS(name string) (os.Os, error) {
	switch name {
	case "centos7":
		return new(centos7.Centos7), nil
	}
	return nil, fmt.Errorf("unsupported os `%s`", name)
}

func NewPackages(name string) (packages.Packages, error) {
	switch name {
	case "yum":
		return new(yum.Yum), nil
	}
	return nil, fmt.Errorf("unsupported package manager `%s`", name)
}

func NewServices(name string) (services.Services, error) {
	switch name {
	case "systemd":
		return new(systemd.Systemd), nil
	}
	return nil, fmt.Errorf("unsupported service manager `%s`", name)
}
