package packages

type Driver interface {
	Setup(c *DriverConfig) error
	InstallPackage(p *Package) error
}

type Package struct {
	Name string
	Removed bool
}

type DriverConfig struct {
	Driver string
	AdditionalSources []string
}