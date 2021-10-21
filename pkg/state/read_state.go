package state

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func ReadServerState(path string, serverName string) (*State, error) {
	var spec State
	if err := hclsimple.DecodeFile(path, nil, &spec); err != nil {
		return nil, fmt.Errorf("can't decode `%s` hcl: %w", path, err)
	}

	for _, serv := range spec.Role.Servers {
		if serv.Name != serverName {
			continue
		}
		spec.Role.Apply(serv)
	}

	spec.Role.Servers = nil

	return &spec, nil
}
