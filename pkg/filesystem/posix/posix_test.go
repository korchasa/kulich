package posix_test

import (
	"github.com/korchasa/kulich/pkg/filesystem"
	"github.com/korchasa/kulich/pkg/filesystem/posix"
	"testing"
)

func TestPosixServices_ImplementInterface(t *testing.T) {
	var _ filesystem.Filesystem = (*posix.Posix)(nil)
}
