package spec

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/slice_diff"
	"strings"
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

func (b Block) Identifier() string {
	return b.Name
}

func (b Block) EqualityHash() string {
	parts := []string{b.Name}
	for _, o := range b.OsOptions {
		parts = append(parts, o.EqualityHash())
	}
	for _, u := range b.Users {
		parts = append(parts, u.EqualityHash())
	}
	for _, p := range b.Packages {
		parts = append(parts, p.EqualityHash())
	}
	for _, d := range b.Directories {
		parts = append(parts, d.EqualityHash())
	}
	for _, f := range b.Files {
		parts = append(parts, f.EqualityHash())
	}
	for _, s := range b.Services {
		parts = append(parts, s.EqualityHash())
	}
	for _, fr := range b.FirewallRules {
		parts = append(parts, fr.EqualityHash())
	}
	return strings.Join(parts, "|")
}

func (b Block) Diff(to Block) (bd BlockDiff, err error) {
	bd.Name = to.Name

	changed, removed, err := slice_diff.SliceDiff(b.OsOptions, to.OsOptions)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for os options: %w", err)
	}
	for _, v := range changed {
		bd.OsOptions.Changed = append(bd.OsOptions.Changed, v.(OsOption))
	}
	for _, v := range removed {
		bd.OsOptions.Removed = append(bd.OsOptions.Removed, v.(OsOption))
	}

	changed, removed, err = slice_diff.SliceDiff(b.Users, to.Users)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for users: %w", err)
	}
	for _, v := range changed {
		bd.Users.Changed = append(bd.Users.Changed, v.(User))
	}
	for _, v := range removed {
		bd.Users.Removed = append(bd.Users.Removed, v.(User))
	}

	changed, removed, err = slice_diff.SliceDiff(b.Packages, to.Packages)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for packages: %w", err)
	}
	for _, v := range changed {
		bd.Packages.Changed = append(bd.Packages.Changed, v.(Package))
	}
	for _, v := range removed {
		bd.Packages.Removed = append(bd.Packages.Removed, v.(Package))
	}

	changed, removed, err = slice_diff.SliceDiff(b.Directories, to.Directories)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for directories: %w", err)
	}
	for _, v := range changed {
		bd.Directories.Changed = append(bd.Directories.Changed, v.(Directory))
	}
	for _, v := range removed {
		bd.Directories.Removed = append(bd.Directories.Removed, v.(Directory))
	}

	changed, removed, err = slice_diff.SliceDiff(b.Files, to.Files)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for files: %w", err)
	}
	for _, v := range changed {
		bd.Files.Changed = append(bd.Files.Changed, v.(File))
	}
	for _, v := range removed {
		bd.Files.Removed = append(bd.Files.Removed, v.(File))
	}

	changed, removed, err = slice_diff.SliceDiff(b.Services, to.Services)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for services: %w", err)
	}
	for _, v := range changed {
		bd.Services.Changed = append(bd.Services.Changed, v.(Service))
	}
	for _, v := range removed {
		bd.Services.Removed = append(bd.Services.Removed, v.(Service))
	}

	changed, removed, err = slice_diff.SliceDiff(b.FirewallRules, to.FirewallRules)
	if err != nil {
		return bd, fmt.Errorf("can't build slice_diff for firewall rules: %w", err)
	}
	for _, v := range changed {
		bd.FirewallRules.Changed = append(bd.FirewallRules.Changed, v.(FirewallRule))
	}
	for _, v := range removed {
		bd.FirewallRules.Removed = append(bd.FirewallRules.Removed, v.(FirewallRule))
	}

	return bd, nil
}

type BlocksDiff struct {
	Changed []BlockDiff
	Removed []Block
}

type BlockDiff struct {
	Name          string
	OsOptions     OsOptionsDiff
	Users         UsersDiff
	Packages      PackagesDiff
	Directories   DirectoriesDiff
	Files         FilesDiff
	Services      ServicesDiff
	FirewallRules FirewallRulesDiff
}
