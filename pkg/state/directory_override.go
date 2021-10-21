package state

import "io/fs"

type DirectoryOverride struct {
	Path        string       `hcl:"path,label"`
	User        *string      `hcl:"user,optional"`
	Group       *string      `hcl:"group,optional"`
	Permissions *fs.FileMode `hcl:"permissions,optional"`
}

func (o *DirectoryOverride) NewDirectory() Directory {
	return Directory{
		Path:        o.Path,
		User:        *o.User,
		Group:       *o.Group,
		Permissions: *o.Permissions,
	}
}
