package centos7_test

import (
	"github.com/korchasa/kulich/pkg/os"
	"github.com/korchasa/kulich/pkg/os/centos7"
	"github.com/korchasa/kulich/pkg/sysshell"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os/exec"
	"testing"
)

type Centos7TestSuite struct {
	suite.Suite
}

func (suite *Centos7TestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestCentos7TestSuite(t *testing.T) {
	suite.Run(t, new(Centos7TestSuite))
}

func (suite *Centos7TestSuite) TestImplementInterface() {
	var _ os.Os = (*centos7.Centos7)(nil)
}

func (suite *Centos7TestSuite) TestSystemd_AddUser() {
	sh := new(sysshell.Mock)
	ops := new(centos7.Centos7)
	assert.NoError(suite.T(), ops.Config(false, sh))

	sh.On("Exec", exec.Command("id", "-u", "user1")).Return(&sysshell.Result{Exit: 1}, nil)
	sh.On("SafeExec", "adduser user1").Return([]string{}, nil)

	err := ops.AddUser(&os.User{Name: "user1"})
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}

func (suite *Centos7TestSuite) TestSystemd_AddUser_System() {
	sh := new(sysshell.Mock)
	ops := new(centos7.Centos7)
	assert.NoError(suite.T(), ops.Config(false, sh))

	sh.On("Exec", exec.Command("id", "-u", "usersys1")).Return(&sysshell.Result{Exit: 1}, nil)
	sh.On("SafeExec", "adduser --system usersys1").Return([]string{}, nil)

	err := ops.AddUser(&os.User{Name: "usersys1", System: true})
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}
