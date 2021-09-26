package packages

type Package struct {
	Name string
	Removed bool
}

type PackageDriverConfig struct {
	Driver string
	AdditionalSources []string
}

type PackageDriver interface {
	Setup(c *PackageDriverConfig) error
	InstallPackage(p *Package) error
}