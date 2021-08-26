package ruchki

import "time"

type User struct {
	Name string
	Shell string
	Home string
}

type OsDriverConfig struct {
	Hostname string
	Selinux string
	Timezone time.Location
	Envs map[string]string
}

type OsDriver interface {
	Setup(c *OsDriverConfig) error
}