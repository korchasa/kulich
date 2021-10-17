package filesystem

import (
	"github.com/korchasa/kulich/pkg/state"
)

type Filesystem interface {
	Config(dryRun bool, opts ...*state.Option) error
	FirstRun() error
	BeforeRun() error
	AddFile(f *state.File) error
	RemoveFile(f *state.File) error
	AddDir(dir *state.Directory) error
	RemoveDir(dir *state.Directory) error
	AfterRun() error
}
