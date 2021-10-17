package state

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileValidate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *File
	}{
		{
			"file path not specified",
			&File{Path: "", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file source not specified",
			&File{Path: ".tmp/dst.txt", From: "", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file user not specified",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", Group: "nobody", Permissions: 0o755},
		},
		{
			"file group not specified",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Permissions: 0o755},
		},
		{
			"file permissions not specified",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody"},
		},
		{
			"can't find file user: user: unknown user unknown",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find file group: group: unknown group unknown",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
		{
			"can't copy file from `./src.txt`: stat ./src.txt: no such file or directory",
			&File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}
