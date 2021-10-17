package posix_test

import (
	"github.com/korchasa/kulich/pkg/sysshell"
	"github.com/korchasa/kulich/pkg/sysshell/posix"
	"testing"
)

func TestPosixSysshell_ImplementInterface(t *testing.T) {
	var _ sysshell.Sysshell = (*posix.Posix)(nil)
}
