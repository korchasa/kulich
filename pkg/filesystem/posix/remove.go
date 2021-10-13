package posix

import (
	"github.com/korchasa/ruchki/pkg/filesystem"
	"os"
)

func (fs *Posix) RemoveFile(f *filesystem.File) error {
	return os.RemoveAll(f.Path)
}

func (fs *Posix) RemoveDir(d *filesystem.Directory) error {
	return os.RemoveAll(d.Path)
}
