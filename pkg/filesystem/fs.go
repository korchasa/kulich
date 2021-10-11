package filesystem

import (
	"fmt"
	"io/fs"
	"strings"
)

type DriverConfig struct {
	Driver  string
	TempDir string
	DryRun  bool
}

type Driver interface {
	Setup(conf *DriverConfig) error
	AddFile(f *File) error
	AddDir(dir *Directory) error
	Delete(path string) error
}

type File struct {
	Path         string
	From         string
	IsTemplate   bool
	TemplateVars interface{}
	IsCompressed bool
	User         string
	Group        string
	Permissions  fs.FileMode
	Hash         string
}

func (f *File) Validate() error {
	if f.Path == "" {
		return fmt.Errorf("file path not specified")
	}
	if f.From == "" {
		return fmt.Errorf("file source not specified")
	}
	if f.User == "" {
		return fmt.Errorf("file user not specified")
	}
	if f.Group == "" {
		return fmt.Errorf("file group not specified")
	}
	if f.Permissions == 0 {
		return fmt.Errorf("file permissions not specified")
	}

	return nil
}

func (f *File) Diffs(a *File) (diffs []string) {
	if f.Path != a.Path {
		diffs = append(diffs, fmt.Sprintf("path: %s != %s", f.Path, a.Path))
	}
	if f.Permissions != a.Permissions {
		diffs = append(diffs, fmt.Sprintf("permisssions: %s != %s", f.Permissions, a.Permissions))
	}
	if f.User != a.User {
		diffs = append(diffs, fmt.Sprintf("user: %s != %s", f.User, a.User))
	}
	if f.Group != a.Group {
		diffs = append(diffs, fmt.Sprintf("group: %s != %s", f.Group, a.Group))
	}
	if f.Hash != a.Hash {
		diffs = append(diffs, fmt.Sprintf("content hash: %s != %s", f.Hash, a.Hash))
	}
	return diffs
}

func (f *File) String() string {
	var sb []string
	sb = append(sb, f.Path)
	sb = append(sb, fmt.Sprintf("from=%s", f.From))
	sb = append(sb, fmt.Sprintf("is_template=%t", f.IsTemplate))
	if f.TemplateVars != nil {
		sb = append(sb, fmt.Sprintf("template_vars=%+v", f.TemplateVars))
	}
	sb = append(sb, fmt.Sprintf("is_compressed=%t", f.IsCompressed))
	sb = append(sb, fmt.Sprintf("user=%s", f.User))
	sb = append(sb, fmt.Sprintf("group=%s", f.Group))
	sb = append(sb, fmt.Sprintf("permissions=%s", f.Permissions))

	return strings.Join(sb, " ")
}

type Directory struct {
	Path        string
	User        string
	Group       string
	Permissions fs.FileMode
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
