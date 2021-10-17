package yum_test

import (
	"github.com/korchasa/kulich/pkg/packages/yum"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestSystemd_Add_Install(t *testing.T) {
	sh := new(sysshell.Mock)
	ym := new(yum.Yum)
	assert.NoError(t, ym.Config(false, sh))

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

	err := ym.Add(&state.Package{Name: "example"})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}

func TestSystemd_Add_AlreadyInstalled(t *testing.T) {
	sh := new(sysshell.Mock)
	ym := new(yum.Yum)
	assert.NoError(t, ym.Config(false, sh))

	sh.
		On("Exec", &exec.Cmd{
			Path: "/usr/bin/yum",
			Args: []string{"yum", "list", "installed", "example", "--assumeyes"},
		}).
		Return(&sysshell.Result{Exit: 0}, nil)

	err := ym.Add(&state.Package{Name: "example"})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}

func TestSystemd_Remove(t *testing.T) {
	sh := new(sysshell.Mock)
	ym := new(yum.Yum)
	assert.NoError(t, ym.Config(false, sh))

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

	err := ym.Remove(&state.Package{Name: "example"})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}
