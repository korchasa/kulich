package services

type Services interface {
	FirstRun() error
	BeforeRun() error
	Add(s *Service) error
	Remove(s *Service) error
	AfterRun() error
}

type Service struct {
	Name            string
	Disabled        bool
	RestartOnChange []Watcher
}

type Watcher struct {
	Path string
	Hash string
}

type Config struct {
	DryRun bool
}
