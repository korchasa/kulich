package state

import (
	"fmt"
	"io/fs"
)

type Directory struct {
	Path        string      `hcl:"path,label"`
	User        string      `hcl:"user"`
	Group       string      `hcl:"group"`
	Permissions fs.FileMode `hcl:"permissions"`
}

func (d *Directory) Validate() error {
	if d.Path == "" {
		return fmt.Errorf("directory path is empty")
	}
	if d.User == "" {
		return fmt.Errorf("directory user is empty")
	}
	if d.Group == "" {
		return fmt.Errorf("directory group is empty")
	}
	if d.Permissions == 0 {
		return fmt.Errorf("directory permissions is empty")
	}

	return nil
}
