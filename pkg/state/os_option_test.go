package state_test

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions_Apply(t *testing.T) {
	var opts state.Options = []state.Option{
		{Type: "foo", Name: "bar", Value: "42"},
	}
	opts.Apply([]state.OptionOverride{
		{Type: "foo", Name: "bar", Value: state.StringRef("43")},
		{Type: "foo", Name: "bar"},
		{Type: "foo", Name: "bar2", Value: state.StringRef("44")},
		{Type: "foo3", Name: "bar3", Value: state.StringRef("45")},
	})
	var expected state.Options = []state.Option{
		{Type: "foo", Name: "bar", Value: "43"},
		{Type: "foo", Name: "bar2", Value: "44"},
		{Type: "foo3", Name: "bar3", Value: "45"},
	}
	assert.Equal(t, expected, opts)
}
