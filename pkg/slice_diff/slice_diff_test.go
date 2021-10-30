package slice_diff_test

import (
	"crypto/sha256"
	"fmt"
	"github.com/korchasa/kulich/pkg/slice_diff"
	"github.com/stretchr/testify/assert"
	"testing"
)

type comp struct {
	Name  string
	Value string
}

func (c comp) Identifier() string {
	return fmt.Sprintf("name=%s", c.Name)
}

func (c comp) EqualityHash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("id=%s,value=%s", c.Identifier(), c.Value)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TestDiff(t *testing.T) {
	from := []comp{
		{Name: "a", Value: "1"},
		{Name: "b", Value: "2"},
		{Name: "b1", Value: "2"},
		{Name: "c", Value: "3"},
		{Name: "d", Value: "4"},
	}
	to := []comp{
		{Name: "a", Value: "1"},
		{Name: "b", Value: "22"},
		{Name: "b1", Value: "2"},
		{Name: "c", Value: "3"},
		{Name: "e", Value: "5"},
	}
	changed, removed, err := slice_diff.SliceDiff(from, to)
	assert.NoError(t, err)
	assert.Len(t, changed, 2)
	assert.IsType(t, changed[0], comp{})
	assert.Equal(t, comp{Name: "b", Value: "22"}, changed[0])
	assert.IsType(t, changed[1], comp{})
	assert.Equal(t, comp{Name: "e", Value: "5"}, changed[1])
	assert.Len(t, removed, 1)
	assert.IsType(t, removed[0], comp{})
	assert.Equal(t, comp{Name: "d", Value: "4"}, removed[0])
}
