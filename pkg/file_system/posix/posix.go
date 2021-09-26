package posix

import (
	"github.com/korchasa/ruchki/pkg/file_system"
)

type Posix struct {
	conf *file_system.FsDriverConfig
	dryRun bool
}

func (fs *Posix) Setup(c *file_system.FsDriverConfig, dryRun bool) error {
	fs.conf = c
	fs.dryRun = dryRun
	return nil
}
