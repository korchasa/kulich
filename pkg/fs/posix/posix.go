package posix

import (
	"github.com/korchasa/ruchki/pkg/fs"
)

type Posix struct {
	conf *fs.DriverConfig
}

func NewPosix(conf *fs.DriverConfig) *Posix {
	return &Posix{conf: conf}
}

func (fs *Posix) Setup(c *fs.DriverConfig) error {
	fs.conf = c
	return nil
}
