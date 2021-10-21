package state

import "io/fs"

type FileOverride struct {
	Path         string            `hcl:"path,label"`
	From         *string           `hcl:"from,optional"`
	IsTemplate   *bool             `hcl:"template,optional"`
	TemplateVars map[string]string `hcl:"template_vars,optional"`
	IsCompressed *bool             `hcl:"compressed,optional"`
	User         *string           `hcl:"user,optional"`
	Group        *string           `hcl:"group,optional"`
	Permissions  *fs.FileMode      `hcl:"permissions,optional"`
	Hash         *string           `hcl:"hash,optional"`
}

func (o *FileOverride) NewFile() File {
	return File{
		Path:         o.Path,
		From:         StringDeref(o.From),
		IsTemplate:   BoolDeref(o.IsTemplate),
		TemplateVars: o.TemplateVars,
		IsCompressed: BoolDeref(o.IsCompressed),
		User:         StringDeref(o.User),
		Group:        StringDeref(o.Group),
		Permissions:  FileModeDeref(o.Permissions),
		Hash:         StringDeref(o.Hash),
	}
}
