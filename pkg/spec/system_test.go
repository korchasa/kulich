package spec_test

import (
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystem_Diff(t *testing.T) {
	from := spec.System{
		OsOptions: []spec.OsOption{
			{Type: "t1", Name: "n1", Value: "v1"},
			{Type: "t2", Name: "n2", Value: "v2"},
			{Type: "t3", Name: "n3", Value: "v3"},
		},
		Users: []spec.User{
			{Name: "n1", System: false},
			{Name: "n2", System: false},
			{Name: "n3", System: false},
		},
		Packages: []spec.Package{
			{Name: "n1", Removed: false},
			{Name: "n2", Removed: false},
			{Name: "n3", Removed: false},
		},
		Directories: []spec.Directory{
			{Path: "p1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", User: "u2", Group: "g2", Permissions: 602},
			{Path: "p3", User: "u3", Group: "g3", Permissions: 603},
		},
		Files: []spec.File{
			{Path: "p1", From: "f1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602},
			{Path: "p3", From: "f3", User: "u3", Group: "g3", Permissions: 603},
		},
		Services: []spec.Service{
			{Name: "n1", Disabled: false},
			{Name: "n2", Disabled: false},
			{Name: "n3", Disabled: false},
		},
		FirewallRules: []spec.FirewallRule{
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
	to := spec.System{
		OsOptions: []spec.OsOption{
			{Type: "t1", Name: "n1", Value: "v1"},
			{Type: "t2", Name: "n2", Value: "v22"},
			{Type: "t4", Name: "n4", Value: "v4"},
		},
		Users: []spec.User{
			{Name: "n1", System: false},
			{Name: "n2", System: true},
			{Name: "n4", System: false},
		},
		Packages: []spec.Package{
			{Name: "n1", Removed: false},
			{Name: "n2", Removed: true},
			{Name: "n4", Removed: false},
		},
		Directories: []spec.Directory{
			{Path: "p1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", User: "u2", Group: "g2", Permissions: 666},
			{Path: "p4", User: "u4", Group: "g4", Permissions: 604},
		},
		Files: []spec.File{
			{Path: "p1", From: "f1", User: "u1", Group: "g1", Permissions: 601},
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602, Hash: "h2"},
			{Path: "p4", From: "f4", User: "u4", Group: "g4", Permissions: 604},
		},
		Services: []spec.Service{
			{Name: "n1", Disabled: false},
			{Name: "n2", Disabled: true},
			{Name: "n4", Disabled: true},
		},
		FirewallRules: []spec.FirewallRule{
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

	diff, err := from.Block().Diff(to.Block())
	assert.NoError(t, err)
	assert.Equal(t, spec.OsOptionsDiff{
		Changed: []spec.OsOption{
			{Type: "t2", Name: "n2", Value: "v22"},
			{Type: "t4", Name: "n4", Value: "v4"},
		},
		Removed: []spec.OsOption{
			{Type: "t3", Name: "n3", Value: "v3"},
		},
	}, diff.OsOptions)
	assert.Equal(t, spec.UsersDiff{
		Changed: []spec.User{
			{Name: "n2", System: true},
			{Name: "n4", System: false},
		},
		Removed: []spec.User{
			{Name: "n3", System: false},
		},
	}, diff.Users)
	assert.Equal(t, spec.PackagesDiff{
		Changed: []spec.Package{
			{Name: "n2", Removed: true},
			{Name: "n4", Removed: false},
		},
		Removed: []spec.Package{
			{Name: "n3", Removed: false},
		},
	}, diff.Packages)
	assert.Equal(t, spec.DirectoriesDiff{
		Changed: []spec.Directory{
			{Path: "p2", User: "u2", Group: "g2", Permissions: 666},
			{Path: "p4", User: "u4", Group: "g4", Permissions: 604},
		},
		Removed: []spec.Directory{
			{Path: "p3", User: "u3", Group: "g3", Permissions: 603},
		},
	}, diff.Directories)
	assert.Equal(t, spec.FilesDiff{
		Changed: []spec.File{
			{Path: "p2", From: "f2", User: "u2", Group: "g2", Permissions: 602, Hash: "h2"},
			{Path: "p4", From: "f4", User: "u4", Group: "g4", Permissions: 604},
		},
		Removed: []spec.File{
			{Path: "p3", From: "f3", User: "u3", Group: "g3", Permissions: 603},
		},
	}, diff.Files)
	assert.Equal(t, spec.ServicesDiff{
		Changed: []spec.Service{
			{Name: "n2", Disabled: true},
			{Name: "n4", Disabled: true},
		},
		Removed: []spec.Service{
			{Name: "n3", Disabled: false},
		},
	}, diff.Services)
	assert.Equal(t, spec.FirewallRulesDiff{
		Changed: []spec.FirewallRule{
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
		Removed: []spec.FirewallRule{
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
