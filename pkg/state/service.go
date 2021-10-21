package state

type Service struct {
	Name            string    `hcl:"name,label"`
	Disabled        bool      `hcl:"disabled,optional"`
	RestartOnChange []Watcher `hcl:"watch,optional"`
}

func (s *Service) Apply(se ServiceOverride) bool {
	if s.Name != se.Name {
		return false
	}
	if se.Disabled != nil {
		s.Disabled = *se.Disabled
	}
	if se.RestartOnChange != nil {
		s.RestartOnChange = se.RestartOnChange
	}
	return true
}

type ServiceOverride struct {
	Name            string    `hcl:"name,label"`
	Disabled        *bool     `hcl:"disabled"`
	RestartOnChange []Watcher `hcl:"watch"`
}

func (o *ServiceOverride) NewService() Service {
	return Service{
		Name:            o.Name,
		Disabled:        *o.Disabled,
		RestartOnChange: o.RestartOnChange,
	}
}

type Watcher struct {
	Path string `hcl:"path"`
	Hash string `hcl:"hash"`
}
