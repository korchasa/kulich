package spec

type Application struct {
	Name          string         `hcl:"name,label"`
	OsOptions     []OsOption     `hcl:"os_option,block"`
	Users         []User         `hcl:"user,block"`
	Packages      []Package      `hcl:"package,block"`
	Directories   []Directory    `hcl:"directory,block"`
	Files         []File         `hcl:"file,block"`
	Services      []Service      `hcl:"service,block"`
	FirewallRules []FirewallRule `hcl:"firewall,block"`
}

func (a Application) Block() Block {
	return Block{
		Name:          a.Name,
		OsOptions:     a.OsOptions,
		Users:         a.Users,
		Packages:      a.Packages,
		Directories:   a.Directories,
		Files:         a.Files,
		Services:      a.Services,
		FirewallRules: a.FirewallRules,
	}
}
