package centos7_test

import (
	"github.com/korchasa/kulich/pkg/os"
	"github.com/korchasa/kulich/pkg/os/centos7"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestImplementInterface(t *testing.T) {
	var _ os.Os = (*centos7.Centos7)(nil)
}

func TestSystemd_AddUser(t *testing.T) {
	sh := new(sysshell.Mock)
	ops := new(centos7.Centos7)
	assert.NoError(t, ops.Config(false, sh))

	sh.On("Exec", exec.Command("id", "-u", "user1")).Return(&sysshell.Result{Exit: 1}, nil)
	sh.On("SafeExec", "adduser user1").Return([]string{}, nil)

	err := ops.AddUser(&state.User{Name: "user1"})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}

func TestSystemd_AddUser_System(t *testing.T) {
	sh := new(sysshell.Mock)
	ops := new(centos7.Centos7)
	assert.NoError(t, ops.Config(false, sh))

	sh.On("Exec", exec.Command("id", "-u", "usersys1")).Return(&sysshell.Result{Exit: 1}, nil)
	sh.On("SafeExec", "adduser --system usersys1").Return([]string{}, nil)

	err := ops.AddUser(&state.User{Name: "usersys1", System: true})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}
