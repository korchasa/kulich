package posix_test

import (
	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/korchasa/ruchki/pkg/filesystem/posix"
	"testing"
)

func TestPosixServices_ImplementInterface(t *testing.T) {
	var _ filesystem.Filesystem = (*posix.Posix)(nil)
}
