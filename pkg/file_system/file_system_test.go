package file_system

import (
	"fmt"
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
			&File{Path: "", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0755},
		},
		{
			"file source not specified",
			&File{Path: "./sandbox/dst.txt", From: "", User: "nobody", Group: "nobody", Permissions: 0755},
		},
		{
			"file user not specified",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", Group: "nobody", Permissions: 0755},
		},
		{
			"file group not specified",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", User: "nobody", Permissions: 0755},
		},
		{
			"file permissions not specified",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody"},
		},
		{
			"can't find file user: user: unknown user unknown",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", User: "unknown", Group: "nobody", Permissions: 0755},
		},
		{
			"can't find file group: group: unknown group unknown",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", User: "nobody", Group: "unknown", Permissions: 0755},
		},
		{
			"can't copy file from `./src.txt`: stat ./src.txt: no such file or directory",
			&File{Path: "./sandbox/dst.txt", From: "./src.txt", User: "nobody", Group: "nobody", Permissions: 0755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, fmt.Errorf(tt.errMsg), tt.spec.Validate())
		})
	}
}

func TestDirValidate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *Directory
	}{
		{
			"directory path is empty",
			&Directory{User: "nobody", Group: "nobody", Permissions: 0755},
		},
		{
			"directory user is empty",
			&Directory{Path: "./test", Group: "nobody", Permissions: 0755},
		},
		{
			"directory group is empty",
			&Directory{Path: "./test", User: "nobody", Permissions: 0755},
		},
		{
			"directory permissions is empty",
			&Directory{Path: "./test", User: "nobody", Group: "nobody"},
		},
		{
			"can't find directory user: user: unknown user unknown",
			&Directory{Path: "./test", User: "unknown", Group: "nobody", Permissions: 0755},
		},
		{
			"can't find directory group: group: unknown group unknown",
			&Directory{Path: "./test", User: "nobody", Group: "unknown", Permissions: 0755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, fmt.Errorf(tt.errMsg), tt.spec.Validate())
		})
	}
}
