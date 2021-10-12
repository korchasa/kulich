package sysshell

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"os/exec"
)

type Sysshell interface {
	Exec(cmd *exec.Cmd) (*Result, error)
	SafeExec(command string) ([]string, error)
	SafeExecf(command string, args ...interface{}) ([]string, error)
}

type Result struct {
	Exit   int
	Stdout []string
	Stderr []string
}

type Mock struct {
	mock.Mock
}

func (m *Mock) Exec(cmd *exec.Cmd) (*Result, error) {
	args := m.Called(cmd)
	return args.Get(0).(*Result), args.Error(1)
}

func (m *Mock) SafeExec(command string) ([]string, error) {
	as := m.Called(command)
	return as.Get(0).([]string), as.Error(1)
}

func (m *Mock) SafeExecf(command string, args ...interface{}) ([]string, error) {
	return m.SafeExec(fmt.Sprintf(command, args...))
}
