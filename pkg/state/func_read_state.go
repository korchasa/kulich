package state

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func ReadServerState(path string, serverName string) (*Root, error) {
	var spec Root
	if err := hclsimple.DecodeFile(path, nil, &spec); err != nil {
		return nil, fmt.Errorf("can't decode `%s` hcl: %w", path, err)
	}

	for _, serv := range spec.Servers {
		if serv.Name != serverName {
			continue
		}
		spec.State.Apply(serv)
	}

	spec.Servers = nil

	return &spec, nil
}
