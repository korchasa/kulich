package packages

type Packages interface {
	FirstRun() error
	BeforeRun() error
	Add(name string) error
	Remove(name string) error
	AfterRun() error
}

type Config struct {
	DryRun bool
}
