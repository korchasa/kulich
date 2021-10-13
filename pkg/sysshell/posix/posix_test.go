package posix_test

import (
	"github.com/korchasa/ruchki/pkg/sysshell"
	"github.com/korchasa/ruchki/pkg/sysshell/posix"
	"testing"
)

func TestPosixSysshell_ImplementInterface(t *testing.T) {
	var _ sysshell.Sysshell = (*posix.Posix)(nil)
}
