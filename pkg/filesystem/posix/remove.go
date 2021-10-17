package posix

import (
	"github.com/korchasa/kulich/pkg/state"
	"os"
)

func (fs *Posix) RemoveFile(f *state.File) error {
	return os.RemoveAll(f.Path)
}

func (fs *Posix) RemoveDir(d *state.Directory) error {
	return os.RemoveAll(d.Path)
}
