package posix

import (
	"github.com/korchasa/kulich/pkg/spec"
	"os"
)

func (fs *Posix) RemoveFile(f *spec.File) error {
	return os.RemoveAll(f.Path)
}

func (fs *Posix) RemoveDir(d *spec.Directory) error {
	return os.RemoveAll(d.Path)
}
