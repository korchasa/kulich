package posix

import "os"

func (fs *Posix) Remove(path string) error {
	return os.RemoveAll(path)
}