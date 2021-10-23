package state

type Service struct {
	Name            string    `hcl:"name,label"`
	Disabled        bool      `hcl:"disabled,optional"`
	RestartOnChange []Watcher `hcl:"watch,optional"`
}

type Watcher struct {
	Path string `hcl:"path"`
	Hash string `hcl:"hash"`
}
