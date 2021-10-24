package spec_test

import (
	"errors"
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_Validate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *spec.File
	}{
		{
			"file path not specified",
			&spec.File{Path: "", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file source not specified",
			&spec.File{Path: ".tmp/dst.txt", From: "", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file user not specified",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", Group: "nobody", Permissions: 0o755},
		},
		{
			"file group not specified",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Permissions: 0o755},
		},
		{
			"file permissions not specified",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody"},
		},
		{
			"can't find file user: user: unknown user unknown",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find file group: group: unknown group unknown",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
		{
			"can't copy file from `./src.txt`: stat ./src.txt: no such file or directory",
			&spec.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}
