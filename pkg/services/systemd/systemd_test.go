package systemd_test

import (
	"github.com/korchasa/kulich/pkg/services/systemd"
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/korchasa/kulich/pkg/sysshell"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemd_Add_NotExists(t *testing.T) {
	service := "example"
	sh := new(sysshell.Mock)
	sys := new(systemd.Systemd)
	assert.NoError(t, sys.Config(false, sh))

	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep Root").
		Return([]string{
			"LoadState=not-found",
			"ActiveState=inactive",
			"SubState=dead",
		}, nil)

	err := sys.Add(&spec.Service{
		Name: service,
	})
	assert.Error(t, err, "service `example` doesn't exists")
}

func TestSystemd_Add(t *testing.T) {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep Root").
		Return([]string{
			"LoadState=loaded",
			"ActiveState=inactive",
			"SubState=dead",
			"UnitFileState=disabled",
		}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl enable example.service").
		Return([]string{}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl start example.service").
		Return([]string{}, nil)

	sys := new(systemd.Systemd)
	assert.NoError(t, sys.Config(false, sh))
	err := sys.Add(&spec.Service{
		Name: service,
	})
	assert.NoError(t, err)

	sh.AssertExpectationsInOrder(t)
}

func TestSystemd_Add_DisableService(t *testing.T) {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep Root").
		Return([]string{
			"LoadState=loaded",
			"ActiveState=active",
			"SubState=running",
			"UnitFileState=enabled",
		}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl disable example.service").
		Return([]string{}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl stop example.service").
		Return([]string{}, nil)

	sys := new(systemd.Systemd)
	assert.NoError(t, sys.Config(false, sh))
	err := sys.Add(&spec.Service{
		Name:     service,
		Disabled: true,
	})

	sh.AssertExpectationsInOrder(t)

	assert.NoError(t, err)
}

func TestSystemd_Remove(t *testing.T) {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep Root").
		Return([]string{
			"LoadState=loaded",
			"ActiveState=active",
			"SubState=running",
			"UnitFileState=enabled",
		}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl disable example.service").
		Return([]string{}, nil)
	sh.
		On("SafeExec", "/usr/bin/systemctl stop example.service").
		Return([]string{}, nil)

	sys := new(systemd.Systemd)
	assert.NoError(t, sys.Config(false, sh))
	err := sys.Remove(&spec.Service{
		Name: service,
	})

	sh.AssertExpectationsInOrder(t)

	assert.NoError(t, err)
}
