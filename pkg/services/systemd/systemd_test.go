package systemd_test

import (
	"github.com/korchasa/ruchki/pkg/services"
	"github.com/korchasa/ruchki/pkg/services/systemd"
	"github.com/korchasa/ruchki/pkg/sysshell"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SystemdIntegrationTestSuite struct {
	suite.Suite
}

func (suite *SystemdIntegrationTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestSystemdIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(SystemdIntegrationTestSuite))
}

func (suite *SystemdIntegrationTestSuite) TestSystemd_Add_NotExists() {
	service := "example"
	sh := new(sysshell.Mock)
	sys := systemd.New(&services.Config{}, sh)

	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep State").
		Return([]string{
			"LoadState=not-found",
			"ActiveState=inactive",
			"SubState=dead",
		}, nil)

	err := sys.Add(&services.Service{
		Name:            service,
		RestartOnChange: nil,
	})
	assert.Error(suite.T(), err, "service `example` doesn't exists")
}

func (suite *SystemdIntegrationTestSuite) TestSystemd_Add() {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep State").
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

	sys := systemd.New(&services.Config{}, sh)
	err := sys.Add(&services.Service{
		Name:            service,
		RestartOnChange: nil,
	})
	assert.NoError(suite.T(), err)

	sh.AssertExpectationsInOrder(suite.T())
}

func (suite *SystemdIntegrationTestSuite) TestSystemd_Add_DisableService() {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep State").
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

	sys := systemd.New(&services.Config{}, sh)
	err := sys.Add(&services.Service{
		Name:            service,
		Disabled:        true,
		RestartOnChange: nil,
	})

	sh.AssertExpectationsInOrder(suite.T())

	assert.NoError(suite.T(), err)
}

func (suite *SystemdIntegrationTestSuite) TestSystemd_Remove() {
	service := "example"

	sh := new(sysshell.Mock)
	sh.
		On("SafeExec", "/usr/bin/systemctl show example.service --no-pager | grep State").
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

	sys := systemd.New(&services.Config{}, sh)
	err := sys.Remove(&services.Service{
		Name:            service,
		RestartOnChange: nil,
	})

	sh.AssertExpectationsInOrder(suite.T())

	assert.NoError(suite.T(), err)
}
