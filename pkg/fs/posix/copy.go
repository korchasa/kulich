package posix

import (
	"fmt"
	"io"
	"os"
)

func (fs *Posix) copy(src string, destination *os.File) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("source file `%s` is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
