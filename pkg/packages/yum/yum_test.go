package yum_test

import (
	"context"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/packages/yum"
	"github.com/korchasa/ruchki/pkg/sysshell"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os/exec"
	"testing"
)

type YumTestSuite struct {
	suite.Suite
}

func (suite *YumTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestSystemdTestSuite(t *testing.T) {
	suite.Run(t, new(YumTestSuite))
}

func (suite *YumTestSuite) TestSystemd_Add_Install() {
	sh := new(sysshell.Mock)
	ym := yum.New(&packages.Config{}, sh)

	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "list", "installed", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 1}, nil)
	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "install", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 0}, nil)

	err := ym.Add(context.TODO(), "example")
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}

func (suite *YumTestSuite) TestSystemd_Add_AlreadyInstalled() {
	sh := new(sysshell.Mock)
	ym := yum.New(&packages.Config{}, sh)

	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "list", "installed", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 0}, nil)

	err := ym.Add(context.TODO(), "example")
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}

func (suite *YumTestSuite) TestSystemd_Remove() {
	sh := new(sysshell.Mock)
	ym := yum.New(&packages.Config{}, sh)

	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "list", "installed", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 0}, nil)
	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "remove", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 0}, nil)

	err := ym.Remove(context.TODO(), "example")
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}
