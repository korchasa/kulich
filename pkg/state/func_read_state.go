package state

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func ReadServerState(path string) (*Root, error) {
	var spec Root
	if err := hclsimple.DecodeFile(path, nil, &spec); err != nil {
		return nil, fmt.Errorf("can't decode `%s` hcl: %w", path, err)
	}

	return &spec, nil
}
