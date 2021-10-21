package state_test

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiles_Apply(t *testing.T) {
	var opts state.Files = []state.File{
		{
			Path:         "/foo",
			From:         "src1",
			IsTemplate:   false,
			TemplateVars: nil,
			IsCompressed: false,
			User:         "usr1",
			Group:        "grp1",
			Permissions:  0400,
			Hash:         "h1",
		},
	}
	opts.Apply([]state.FileOverride{
		{Path: "/foo", From: state.StringRef("src2"), IsTemplate: state.BoolRef(true), User: state.StringRef("usr2")},
		{Path: "/foo", From: state.StringRef("src3"), User: state.StringRef("usr3")},
		{Path: "/foo", From: state.StringRef("src4"), Group: state.StringRef("grp2")},
		{Path: "/bar", From: state.StringRef("src1"), Group: state.StringRef("grp1")},
	})
	var expected state.Files = []state.File{
		{
			Path:         "/foo",
			From:         "src4",
			IsTemplate:   true,
			TemplateVars: nil,
			IsCompressed: false,
			User:         "usr3",
			Group:        "grp2",
			Permissions:  0400,
			Hash:         "h1",
		},
		{
			Path:  "/bar",
			From:  "src1",
			Group: "grp1",
		},
	}
	assert.Equal(t, expected, opts)
}
