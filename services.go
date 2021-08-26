package ruchki

type Service struct {
	Name string
	Enabled bool
}

type ServicesDriverConfig struct {
	Driver string
}

type ServicesDriver interface {
	Enable(name string) error
	Disable(name string) error
}