package services

type Driver interface {
	Add(s *Service) error
	Remove(s *Service) error
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

type DriverConfig struct {
	DryRun bool
}
