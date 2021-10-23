package state_test

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystem_Diff(t *testing.T) {
	from := state.System{
		OsOptions: []state.OsOption{
			{Type: "t1", Name: "n1", Value: "v1"},
			{Type: "t2", Name: "n2", Value: "v2"},
			{Type: "t3", Name: "n3", Value: "v3"},
		},
		Users: []state.User{
			{Name: "n1", System: false},
			{Name: "n2", System: false},
			{Name: "n3", System: false},
		},
		Packages: []state.Package{
			{Name: "n1", Removed: false},
			{Name: "n2", Removed: false},
			{Name: "n3", Removed: false},
		},
		Directories: []state.Directory{
			{Path: "p1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", User: "u2", Group: "g2", Permissions: 602},
			{Path: "p3", User: "u3", Group: "g3", Permissions: 603},
		},
		Files: []state.File{
			{Path: "p1", From: "f1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602},
			{Path: "p3", From: "f3", User: "u3", Group: "g3", Permissions: 603},
		},
		Services: []state.Service{
			{Name: "n1", Disabled: false},
			{Name: "n2", Disabled: false},
			{Name: "n3", Disabled: false},
		},
		FirewallRules: []state.FirewallRule{
			{
				Name:     "n1",
				Id:       "i1",
				Ports:    []string{"p11", "p12"},
				Protocol: "p1",
				Targets:  []string{"t11", "t12"},
				IsOutput: false,
			},
			{
				Name:     "n2",
				Id:       "i2",
				Ports:    []string{"p21", "p22"},
				Protocol: "p2",
				Targets:  []string{"t21", "t22"},
				IsOutput: false,
			},
			{
				Name:     "n3",
				Id:       "i3",
				Ports:    []string{"p31", "p32"},
				Protocol: "p3",
				Targets:  []string{"t31", "t32"},
				IsOutput: false,
			},
		},
	}
	to := state.System{
		OsOptions: []state.OsOption{
			{Type: "t1", Name: "n1", Value: "v1"},
			{Type: "t2", Name: "n2", Value: "v22"},
			{Type: "t4", Name: "n4", Value: "v4"},
		},
		Users: []state.User{
			{Name: "n1", System: false},
			{Name: "n2", System: true},
			{Name: "n4", System: false},
		},
		Packages: []state.Package{
			{Name: "n1", Removed: false},
			{Name: "n2", Removed: true},
			{Name: "n4", Removed: false},
		},
		Directories: []state.Directory{
			{Path: "p1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", User: "u2", Group: "g2", Permissions: 666},
			{Path: "p4", User: "u4", Group: "g4", Permissions: 604},
		},
		Files: []state.File{
			{Path: "p1", From: "f1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602, Hash: "h2"},
			{Path: "p4", From: "f4", User: "u4", Group: "g4", Permissions: 604},
		},
		Services: []state.Service{
			{Name: "n1", Disabled: false},
			{Name: "n2", Disabled: true},
			{Name: "n4", Disabled: true},
		},
		FirewallRules: []state.FirewallRule{
			{
				Name:     "n1",
				Id:       "i1",
				Ports:    []string{"p11", "p12"},
				Protocol: "p1",
				Targets:  []string{"t11", "t12"},
				IsOutput: false,
			},
			{
				Name:     "n2",
				Id:       "i2",
				Ports:    []string{"p21", "p22", "p23"},
				Protocol: "p2",
				Targets:  []string{"t21", "t22", "p23"},
				IsOutput: true,
			},
			{
				Name:     "n4",
				Id:       "i4",
				Ports:    []string{"p41", "p42"},
				Protocol: "p4",
				Targets:  []string{"t41", "t42"},
				IsOutput: true,
			},
		},
	}

	diff, err := from.Diff(to)
	assert.NoError(t, err)
	assert.Equal(t, state.OsOptionsDiff{
		Changed: []state.OsOption{
			{Type: "t2", Name: "n2", Value: "v22"},
			{Type: "t4", Name: "n4", Value: "v4"},
		},
		Removed: []state.OsOption{
			{Type: "t3", Name: "n3", Value: "v3"},
		},
	}, diff.OsOptions)
	assert.Equal(t, state.UsersDiff{
		Changed: []state.User{
			{Name: "n2", System: true},
			{Name: "n4", System: false},
		},
		Removed: []state.User{
			{Name: "n3", System: false},
		},
	}, diff.Users)
	assert.Equal(t, state.PackagesDiff{
		Changed: []state.Package{
			{Name: "n2", Removed: true},
			{Name: "n4", Removed: false},
		},
		Removed: []state.Package{
			{Name: "n3", Removed: false},
		},
	}, diff.Packages)
	assert.Equal(t, state.DirectoriesDiff{
		Changed: []state.Directory{
			{Path: "p2", User: "u2", Group: "g2", Permissions: 666},
			{Path: "p4", User: "u4", Group: "g4", Permissions: 604},
		},
		Removed: []state.Directory{
			{Path: "p3", User: "u3", Group: "g3", Permissions: 603},
		},
	}, diff.Directories)
	assert.Equal(t, state.FilesDiff{
		Changed: []state.File{
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602, Hash: "h2"},
			{Path: "p4", From: "f4", User: "u4", Group: "g4", Permissions: 604},
		},
		Removed: []state.File{
			{Path: "p3", From: "f3", User: "u3", Group: "g3", Permissions: 603},
		},
	}, diff.Files)
	assert.Equal(t, state.ServicesDiff{
		Changed: []state.Service{
			{Name: "n2", Disabled: true},
			{Name: "n4", Disabled: true},
		},
		Removed: []state.Service{
			{Name: "n3", Disabled: false},
		},
	}, diff.Services)
	assert.Equal(t, state.FirewallRulesDiff{
		Changed: []state.FirewallRule{
			{
				Name:     "n2",
				Id:       "i2",
				Ports:    []string{"p21", "p22", "p23"},
				Protocol: "p2",
				Targets:  []string{"t21", "t22", "p23"},
				IsOutput: true,
			},
			{
				Name:     "n4",
				Id:       "i4",
				Ports:    []string{"p41", "p42"},
				Protocol: "p4",
				Targets:  []string{"t41", "t42"},
				IsOutput: true,
			},
		},
		Removed: []state.FirewallRule{
			{
				Name:     "n3",
				Id:       "i3",
				Ports:    []string{"p31", "p32"},
				Protocol: "p3",
				Targets:  []string{"t31", "t32"},
				IsOutput: false,
			},
		},
	}, diff.FirewallRules)
}
