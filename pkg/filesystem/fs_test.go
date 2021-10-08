package filesystem_test

import (
	"errors"
	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileValidate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *filesystem.File
	}{
		{
			"file path not specified",
			&filesystem.File{Path: "", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file source not specified",
			&filesystem.File{Path: ".tmp/dst.txt", From: "", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file user not specified",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", Group: "nobody", Permissions: 0o755},
		},
		{
			"file group not specified",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Permissions: 0o755},
		},
		{
			"file permissions not specified",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody"},
		},
		{
			"can't find file user: user: unknown user unknown",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find file group: group: unknown group unknown",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
		{
			"can't copy file from `./src.txt`: stat ./src.txt: no such file or directory",
			&filesystem.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}

func TestDirValidate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *filesystem.Directory
	}{
		{
			"directory path is empty",
			&filesystem.Directory{User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"directory user is empty",
			&filesystem.Directory{Path: "./test", Group: "nobody", Permissions: 0o755},
		},
		{
			"directory group is empty",
			&filesystem.Directory{Path: "./test", User: "nobody", Permissions: 0o755},
		},
		{
			"directory permissions is empty",
			&filesystem.Directory{Path: "./test", User: "nobody", Group: "nobody"},
		},
		{
			"can't find directory user: user: unknown user unknown",
			&filesystem.Directory{Path: "./test", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find directory group: group: unknown group unknown",
			&filesystem.Directory{Path: "./test", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}
