package posix

import (
	"fmt"
	"io"
	"os"
)

func (fs *Posix) copy(dst *os.File, src string) (int64, error) {
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

	nBytes, err := io.Copy(dst, source)

	return nBytes, err
}
