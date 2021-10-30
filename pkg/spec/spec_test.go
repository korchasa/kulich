package spec_test

import (
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpec_Diff(t *testing.T) {
	from := &spec.Spec{
		Name:   "spec",
		Config: spec.Config{},
		System: spec.Block{
			Name: "system",
			OsOptions: []spec.OsOption{
				{Type: "sys_ot1", Name: "sys_on1", Value: "sys_ov1"},
				{Type: "sys_ot2", Name: "sys_on2", Value: "sys_ov2"},
				{Type: "sys_ot3", Name: "sys_on3", Value: "sys_ov3"},
			},
		},
		Applications: []spec.Block{
			{
				Name: "app1",
				OsOptions: []spec.OsOption{
					{Type: "app1_ot1", Name: "app1_on1", Value: "app1_ov1"},
				},
			},
			{
				Name: "app2",
				OsOptions: []spec.OsOption{
					{Type: "app2_ot1", Name: "app2_on1", Value: "app2_ov1"},
				},
			},
			{
				Name: "app3",
				OsOptions: []spec.OsOption{
					{Type: "app3_ot1", Name: "app3_on1", Value: "app3_ov1"},
				},
			},
		},
	}
	to := spec.Spec{
		Name:   "spec",
		Config: spec.Config{},
		System: spec.Block{
			Name: "system",
			OsOptions: []spec.OsOption{
				{Type: "sys_ot1", Name: "sys_on1", Value: "sys_ov1"},
				{Type: "sys_ot2", Name: "sys_on2", Value: "sys_ov2_changed"},
				{Type: "sys_ot4", Name: "sys_on4", Value: "sys_ov4"},
			},
		},
		Applications: []spec.Block{
			{
				Name: "app1",
				OsOptions: []spec.OsOption{
					{Type: "app1_ot1", Name: "app1_on1", Value: "app1_ov1"},
				},
			},
			{
				Name: "app2",
				OsOptions: []spec.OsOption{
					{Type: "app2_ot1", Name: "app2_on1", Value: "app2_ov1_changed"},
				},
			},
			{
				Name: "app4",
				OsOptions: []spec.OsOption{
					{Type: "app4_ot1", Name: "app4_on1", Value: "app4_ov1"},
				},
			},
		},
	}
	diff, err := from.Diff(to)
	assert.NoError(t, err)
	assert.Equal(t, spec.Diff{
		System: spec.BlockDiff{
			Name: "system",
			OsOptions: spec.OsOptionsDiff{
				Changed: []spec.OsOption{
					{Type: "sys_ot2", Name: "sys_on2", Value: "sys_ov2_changed"},
					{Type: "sys_ot4", Name: "sys_on4", Value: "sys_ov4"},
				},
				Removed: []spec.OsOption{
					{Type: "sys_ot3", Name: "sys_on3", Value: "sys_ov3"},
				},
			},
		},
		Applications: spec.BlocksDiff{
			Changed: []spec.BlockDiff{
				{
					Name: "app2",
					OsOptions: spec.OsOptionsDiff{
						Changed: []spec.OsOption{
							{Type: "app2_ot1", Name: "app2_on1", Value: "app2_ov1_changed"},
						},
					},
				},
				{
					Name: "app4",
					OsOptions: spec.OsOptionsDiff{
						Changed: []spec.OsOption{
							{Type: "app4_ot1", Name: "app4_on1", Value: "app4_ov1"},
						},
					},
				},
			},
			Removed: []spec.Block{
				{
					Name: "app3",
					OsOptions: []spec.OsOption{
						{Type: "app3_ot1", Name: "app3_on1", Value: "app3_ov1"},
					},
				},
			},
		},
	}, diff)
}
