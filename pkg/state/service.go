package state

type Service struct {
	Name            string `hcl:"name,label"`
	Disabled        bool   `hcl:"disabled,optional"`
	RestartOnChange []Watcher
}

type Watcher struct {
	Path string
	Hash string
}

type ServiceOverride struct {
	Name            string `hcl:"name,label"`
	Disabled        *bool  `hcl:"disabled"`
	RestartOnChange []Watcher
}
