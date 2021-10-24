package spec

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirValidate(t *testing.T) {
	tests := []struct {
		errMsg string
		spec   *Directory
	}{
		{
			"directory path is empty",
			&Directory{User: "nobody", Group: "nobody", Permissions: 0o755},
		},
		{
			"directory user is empty",
			&Directory{Path: "./test", Group: "nobody", Permissions: 0o755},
		},
		{
			"directory group is empty",
			&Directory{Path: "./test", User: "nobody", Permissions: 0o755},
		},
		{
			"directory permissions is empty",
			&Directory{Path: "./test", User: "nobody", Group: "nobody"},
		},
		{
			"can't find directory user: user: unknown user unknown",
			&Directory{Path: "./test", User: "unknown", Group: "nobody", Permissions: 0o755},
		},
		{
			"can't find directory group: group: unknown group unknown",
			&Directory{Path: "./test", User: "nobody", Group: "unknown", Permissions: 0o755},
		},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			assert.Error(t, errors.New(tt.errMsg), tt.spec.Validate())
		})
	}
}
