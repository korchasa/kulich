package posix

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/config"
	"os"
)

type Posix struct {
	dryRun  bool
	tempDir string
}

func (fs *Posix) Config(dryRun bool, opts ...*config.Option) error {
	fs.dryRun = dryRun
	for _, v := range opts {
		switch v.Type {
		case "temp_dir":
			fs.tempDir = v.Value
		default:
			return fmt.Errorf("unsupported option type `%s`", v.Type)
		}
	}

	return nil
}

func (fs *Posix) FirstRun() error {
	return nil
}

func (fs *Posix) BeforeRun() error {
	return nil
}

func (fs *Posix) AfterRun() error {
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
