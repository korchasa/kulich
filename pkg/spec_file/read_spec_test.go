package spec_file_test

import (
	"github.com/korchasa/kulich/pkg/spec_file"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadSpec(t *testing.T) {
	st, err := spec_file.ReadSpec("./fixtures/full.hcl")
	assert.NoError(t, err)
	assert.Equal(
		t,
		spec_file.Root{
			Spec: spec_file.Spec{
				Name: "nomad-clients",
				Config: spec_file.Config{
					OsDriver:         spec_file.DriverConfig{Name: "centos7"},
					PackagesDriver:   spec_file.DriverConfig{Name: "yum"},
					FilesystemDriver: spec_file.DriverConfig{Name: "posix"},
					ServicesDriver:   spec_file.DriverConfig{Name: "systemctl"},
					FirewallDriver:   spec_file.DriverConfig{Name: "iptables"},
				},
				System: spec_file.System{
					OsOptions: []spec_file.OsOption{
						{Type: "hostnamectl", Name: "hostname", Value: "node1-nomad"},
						{Type: "selinux", Name: "enabled", Value: "false"},
					},
					Users: []spec_file.User{{
						Name:   "alice",
						System: false,
					}},
					Packages: []spec_file.Package{
						{Name: "epel-release", Removed: false},
						{Name: "yum-utils", Removed: false},
						{Name: "unzip", Removed: false},
						{Name: "vim-7.29.0-59.el7", Removed: false},
						{Name: "noderig", Removed: true},
						{Name: "firewalld", Removed: true},
					},
					Files: []spec_file.File{
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
				Applications: []spec_file.Application{
					{
						Name: "consul",
						OsOptions: []spec_file.OsOption{{
							Type:  "env",
							Name:  "CONSUL_HTTP_ADDR",
							Value: "http://127.0.0.1:8500",
						}},
						Users: []spec_file.User{{
							Name:   "consul",
							System: true,
						}},
						Packages: []spec_file.Package{
							{Name: "unbound", Removed: false},
						},
						Directories: []spec_file.Directory{
							{
								Path:        "/var/consul/",
								User:        "consul",
								Group:       "consul",
								Permissions: 700,
							},
						},
						Files: []spec_file.File{
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
						Services: []spec_file.Service{
							{
								Name:     "consul",
								Disabled: false,
							},
						},
						FirewallRules: []spec_file.FirewallRule{
							{
								Name:  "consul",
								Ports: []string{"8300:8302", "8500:8502", "8600"},
								Targets: []string{
									"444.444.444.444/16",
									"555.555.555.555",
									"666.666.666.66",
								},
							},
						},
					},
				},
			},
		},
		*st,
	)
}
