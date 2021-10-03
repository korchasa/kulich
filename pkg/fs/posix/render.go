package posix

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"io"
	"os"
	"text/template"
)

func (fs *Posix) render(f *os.File, vars interface{}) error {
	buf := bytes.NewBuffer(nil)
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("can't return to begin of file: %v", err)
	}
	if _, err := io.Copy(buf, f); err != nil {
		return fmt.Errorf("can't read file content: %v", err)
	}
	t, err := template.New("common").Funcs(sprig.TxtFuncMap()).Parse(buf.String())
	if err != nil {
		return fmt.Errorf("can't parse template: %v", err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("can't return to begin of file: %v", err)
	}
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("can't truncate file: %v", err)
	}
	err = t.Execute(f, vars)
	if err != nil {
		return fmt.Errorf("can't execute template: %v", err)
	}
	return nil
}
