package os

import "time"

type Driver interface {
	Setup(c *DriverConfig) error
}

type User struct {
	Name  string
	Shell string
	Home  string
}

type DriverConfig struct {
	Hostname string
	Selinux  string
	Timezone time.Location
	Envs     map[string]string
}
