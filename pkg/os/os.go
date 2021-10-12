package os

import "time"

type Os interface {
	Setup(c *Config) error
	AddUser(u *User)
}

type User struct {
	Name           string
	Shell          string
	Home           string
	AuthorizedKeys string
	Removed        bool
}

type Config struct {
	Hostname string
	Selinux  string
	Timezone time.Location
	Envs     map[string]string
}
