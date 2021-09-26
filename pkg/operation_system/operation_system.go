package operation_system

import "time"

type User struct {
	Name string
	Shell string
	Home string
}

type DriverConfig struct {
	Hostname string
	Selinux string
	Timezone time.Location
	Envs map[string]string
}

type Driver interface {
	Setup(c *DriverConfig) error
}