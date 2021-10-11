package os

import "time"

type Os interface {
	Setup(c *Config) error
}

type User struct {
	Name  string
	Shell string
	Home  string
}

type Config struct {
	Hostname string
	Selinux  string
	Timezone time.Location
	Envs     map[string]string
}
