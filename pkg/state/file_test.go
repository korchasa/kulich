package state_test

import (
	"errors"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_Validate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *state.File
	}{
		{
			"file path not specified",
			&state.File{Path: "", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file source not specified",
			&state.File{Path: ".tmp/dst.txt", From: "", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"file user not specified",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", Group: "nobody", Permissions: 0o755},
		},
		{
			"file group not specified",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Permissions: 0o755},
		},
		{
			"file permissions not specified",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody"},
		},
		{
			"can't find file user: user: unknown user unknown",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find file group: group: unknown group unknown",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
		{
			"can't copy file from `./src.txt`: stat ./src.txt: no such file or directory",
			&state.File{Path: ".tmp/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}
