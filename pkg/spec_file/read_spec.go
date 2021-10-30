package spec_file

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func ReadSpec(path string) (*Root, error) {
	var spec Root
	if err := hclsimple.DecodeFile(path, nil, &spec); err != nil {
		return nil, fmt.Errorf("can't decode `%s` hcl: %w", path, err)
	}

	return &spec, nil
}
