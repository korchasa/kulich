package ruchki

import (
	"io"
	"io/fs"
	"net/url"
)

type File struct {
	Path string
	Content io.ByteReader
	FromUri url.URL
	FromTemplate string
	Vars interface{}
	Compressed bool
	Owner string
	Group string
	Permissions fs.FileMode
}

type Directory struct {
	Path string
	Owner string
	Group string
	Permissions fs.FileMode
}

type FsDriverConfig struct {
	Driver string
}

type FsDriver interface {
	Setup(c *FsDriverConfig) error
	ApplyFile(f *File) error
	ApplyDir(d *Directory) error
}
