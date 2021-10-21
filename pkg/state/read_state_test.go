package state_test

import (
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadServerState(t *testing.T) {
	st, err := state.ReadServerState("./fixtures/full.hcl", "node1-nomad")
	assert.NoError(t, err)
	assert.Equal(
		t,
		state.State{
			Role: state.Role{
				Name: "nomad-clients",
				Config: state.Config{
					OsDriver:         state.DriverConfig{Name: "centos7"},
					PackagesDriver:   state.DriverConfig{Name: "yum"},
					FilesystemDriver: state.DriverConfig{Name: "posix"},
					ServicesDriver:   state.DriverConfig{Name: "systemctl"},
					FirewallDriver:   state.DriverConfig{Name: "iptables"},
				},
				System: state.System{
					OsOptions: []state.Option{
						{Type: "selinux", Name: "enabled", Value: "false"},
						{Type: "hostnamectl", Name: "hostname", Value: "node1-nomad"},
					},
					Users: []state.User{{
						Name:   "alice",
						System: false,
					}},
					Packages: []state.Package{
						{Name: "epel-release", Removed: false},
						{Name: "yum-utils", Removed: false},
						{Name: "unzip", Removed: false},
						{Name: "vim-7.29.0-59.el7", Removed: false},
						{Name: "noderig", Removed: true},
						{Name: "firewalld", Removed: true},
					},
					Files: []state.File{
						{
							Path:        "/home/korchasa/.ssh/authorized_keys",
							From:        "./korchasa_authorized_keys",
							User:        "alice",
							Group:       "alice",
							Permissions: 600,
						},
						{
							Path:        "/etc/yum.repos.d/docker-ce.repo",
							From:        "./docker-ce.repo",
							User:        "consul",
							Group:       "consul",
							Permissions: 600,
						},
					},
				},
				Applications: []state.Application{
					{
						Name: "consul",
						OsOptions: []state.Option{{
							Type:  "env",
							Name:  "CONSUL_HTTP_ADDR",
							Value: "http://127.0.0.1:8500",
						}},
						Users: []state.User{{
							Name:   "consul",
							System: true,
						}},
						Packages: []state.Package{
							{Name: "unbound", Removed: false},
						},
						Directories: []state.Directory{
							{
								Path:        "/var/consul/",
								User:        "consul",
								Group:       "consul",
								Permissions: 700,
							},
						},
						Files: []state.File{
							{
								Path:         "/usr/local/bin/consul",
								From:         "https://releases.hashicorp.com/consul/1.9.5/consul_1.9.5_linux_amd64.zip",
								IsCompressed: true,
								User:         "consul",
								Group:        "consul",
								Permissions:  600,
							},
							{
								Path:       "/etc/consul.d/config.json",
								From:       "./consul_client_config.json",
								IsTemplate: true,
								TemplateVars: map[string]string{
									"foo": "bar",
									"baz": "42",
								},
								User:        "consul",
								Group:       "consul",
								Permissions: 400,
							},
						},
						Services: []state.Service{
							{
								Name:     "consul",
								Disabled: false,
							},
						},
						FirewallRules: []state.FirewallRule{
							{
								Name:  "consul",
								Ports: []string{"8300:8302", "8500:8502", "8600"},
								Targets: []string{
									"111.111.111.111/16",
									"222.222.222.222",
									"333.333.333.333",
								},
							},
						},
					},
					{
						Name: "some-app",
						Files: []state.File{
							{
								Path:        "/tmp/some-app.txt",
								From:        "./some-app.txt",
								User:        "consul",
								Group:       "consul",
								Permissions: 600,
							},
						},
					},
				},
			},
		},
		*st,
	)
}
