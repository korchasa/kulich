package services

type Driver interface {
	Enable(name string) error
	Disable(name string) error
}

type Service struct {
	Name string
	Enabled bool
}

type DriverConfig struct {
	Driver string
}
