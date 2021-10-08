package posix

import (
	"fmt"
	"github.com/korchasa/ruchki/pkg/filesystem"
	"os"
)

type Posix struct {
	conf *filesystem.DriverConfig
}

func NewPosix(conf *filesystem.DriverConfig) *Posix {
	return &Posix{conf: conf}
}

func (fs *Posix) Setup(c *filesystem.DriverConfig) error {
	fs.conf = c
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("can't get file stat: %w", err)
	}

	return true, nil
}
