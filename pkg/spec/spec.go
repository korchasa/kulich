package spec

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/slice_diff"
)

type Spec struct {
	Name         string
	Config       Config
	System       Block
	Applications []Block
}

func (s Spec) Diff(to Spec) (d Diff, err error) {
	d.System, err = s.System.Diff(to.System)
	if err != nil {
		return d, fmt.Errorf("can't build slice_diff for system block: %w", err)
	}
	chApps, rmApps, err := slice_diff.SliceDiff(s.Applications, to.Applications)
	for _, chAppInt := range chApps {
		chApp, ok := chAppInt.(Block)
		if !ok {
			return d, fmt.Errorf("can't cast changed application to struct")
		}
		for _, toApp := range to.Applications {
			if chApp.Name != toApp.Name {
				continue
			}
			found := false
			for _, fromApp := range s.Applications {
				if chApp.Name != fromApp.Name {
					continue
				}
				diff, err := fromApp.Diff(toApp)
				if err != nil {
					return d, fmt.Errorf("can't build diff for `%s`: %w", chApp.Name, err)
				}
				d.Applications.Changed = append(d.Applications.Changed, diff)
				found = true
			}
			if !found {
				diff, err := Block{}.Diff(toApp)
				if err != nil {
					return d, fmt.Errorf("can't build diff for `%s` from empty app: %w", chApp.Name, err)
				}
				d.Applications.Changed = append(d.Applications.Changed, diff)
			}
		}
	}
	for _, rmAppInt := range rmApps {
		rmApp, ok := rmAppInt.(Block)
		if !ok {
			return d, fmt.Errorf("can't cast removed application to struct")
		}
		d.Applications.Removed = append(d.Applications.Removed, rmApp)
	}

	return d, nil
}

type Diff struct {
	System       BlockDiff
	Applications BlocksDiff
}
