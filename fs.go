package ruchki

import (
	"io"
	"io/fs"
	"net/url"
)

type File struct {
	Path string
	FromUri *url.URL
	FromTemplate *string
	Vars *interface{}
	FromContent *io.ByteReader
	Compressed *bool
	User string
	Group string
	Permissions fs.FileMode
}

type Directory struct {
	Path  string
	User  string
	Group string
	Permissions fs.FileMode
}

type FsDriverConfig struct {
	Driver string
	DryRun bool
}

type FsDriver interface {
	Setup(c *FsDriverConfig) error
	ApplyFile(f *File) error
	ApplyDir(d *Directory) error
}
