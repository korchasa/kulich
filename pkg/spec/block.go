package spec

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/diff"
)

type Block struct {
	Name          string
	OsOptions     []OsOption
	Users         []User
	Packages      []Package
	Directories   []Directory
	Files         []File
	Services      []Service
	FirewallRules []FirewallRule
}

func (s Block) Diff(to Block) (bd BlockDiff, err error) {
	bd.Name = s.Name

	changed, removed, err := diff.Diff(s.OsOptions, to.OsOptions)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for os options: %w", err)
	}
	for _, v := range changed {
		bd.OsOptions.Changed = append(bd.OsOptions.Changed, v.(OsOption))
	}
	for _, v := range removed {
		bd.OsOptions.Removed = append(bd.OsOptions.Removed, v.(OsOption))
	}

	changed, removed, err = diff.Diff(s.Users, to.Users)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for users: %w", err)
	}
	for _, v := range changed {
		bd.Users.Changed = append(bd.Users.Changed, v.(User))
	}
	for _, v := range removed {
		bd.Users.Removed = append(bd.Users.Removed, v.(User))
	}

	changed, removed, err = diff.Diff(s.Packages, to.Packages)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for packages: %w", err)
	}
	for _, v := range changed {
		bd.Packages.Changed = append(bd.Packages.Changed, v.(Package))
	}
	for _, v := range removed {
		bd.Packages.Removed = append(bd.Packages.Removed, v.(Package))
	}

	changed, removed, err = diff.Diff(s.Directories, to.Directories)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for directories: %w", err)
	}
	for _, v := range changed {
		bd.Directories.Changed = append(bd.Directories.Changed, v.(Directory))
	}
	for _, v := range removed {
		bd.Directories.Removed = append(bd.Directories.Removed, v.(Directory))
	}

	changed, removed, err = diff.Diff(s.Files, to.Files)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for files: %w", err)
	}
	for _, v := range changed {
		bd.Files.Changed = append(bd.Files.Changed, v.(File))
	}
	for _, v := range removed {
		bd.Files.Removed = append(bd.Files.Removed, v.(File))
	}

	changed, removed, err = diff.Diff(s.Services, to.Services)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for services: %w", err)
	}
	for _, v := range changed {
		bd.Services.Changed = append(bd.Services.Changed, v.(Service))
	}
	for _, v := range removed {
		bd.Services.Removed = append(bd.Services.Removed, v.(Service))
	}

	changed, removed, err = diff.Diff(s.FirewallRules, to.FirewallRules)
	if err != nil {
		return bd, fmt.Errorf("can't build diff for firewall rules: %w", err)
	}
	for _, v := range changed {
		bd.FirewallRules.Changed = append(bd.FirewallRules.Changed, v.(FirewallRule))
	}
	for _, v := range removed {
		bd.FirewallRules.Removed = append(bd.FirewallRules.Removed, v.(FirewallRule))
	}

	return bd, nil
}
