package state

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	var spec State
	err := hclsimple.DecodeFile("./fixture.hcl", nil, &spec)
	assert.NoError(t, err)
}
