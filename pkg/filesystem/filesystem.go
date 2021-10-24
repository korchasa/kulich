package filesystem

import (
	"github.com/korchasa/kulich/pkg/spec"
)

type Filesystem interface {
	Config(dryRun bool, opts ...*spec.OsOption) error
	FirstRun() error
	BeforeRun() error
	AddFile(f *spec.File) error
	RemoveFile(f *spec.File) error
	AddDir(dir *spec.Directory) error
	RemoveDir(dir *spec.Directory) error
	AfterRun() error
}
