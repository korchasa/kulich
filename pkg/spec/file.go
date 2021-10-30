package spec

import (
	"fmt"
	"io/fs"
	"strings"
)

type File struct {
	Path         string
	From         string
	IsTemplate   bool
	TemplateVars map[string]string
	IsCompressed bool
	User         string
	Group        string
	Permissions  fs.FileMode
	Hash         string
}

func (f File) Identifier() string {
	return f.Path
}

func (f File) EqualityHash() string {
	return fmt.Sprintf("%v", f)
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

func (f *File) String() string {
	var sb []string
	sb = append(sb, f.Path)
	sb = append(sb, fmt.Sprintf("from=%s", f.From))
	sb = append(sb, fmt.Sprintf("is_template=%v", f.IsTemplate))
	if f.TemplateVars != nil {
		sb = append(sb, fmt.Sprintf("template_vars=%+v", f.TemplateVars))
	}
	sb = append(sb, fmt.Sprintf("is_compressed=%v", f.IsCompressed))
	sb = append(sb, fmt.Sprintf("user=%s", f.User))
	sb = append(sb, fmt.Sprintf("group=%s", f.Group))
	sb = append(sb, fmt.Sprintf("permissions=%s", f.Permissions))

	return strings.Join(sb, " ")
}

type Files []File

type FilesDiff struct {
	Changed []File
	Removed []File
}
