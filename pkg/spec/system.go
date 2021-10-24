package spec

type System struct {
	OsOptions     []OsOption     `hcl:"os_option,block"`
	Users         []User         `hcl:"user,block"`
	Packages      []Package      `hcl:"package,block"`
	Directories   []Directory    `hcl:"directory,block"`
	Files         []File         `hcl:"file,block"`
	Services      []Service      `hcl:"service,block"`
	FirewallRules []FirewallRule `hcl:"firewall,block"`
}

func (s System) Block() Block {
	return Block{
		Name:          "system",
		OsOptions:     s.OsOptions,
		Users:         s.Users,
		Packages:      s.Packages,
		Directories:   s.Directories,
		Files:         s.Files,
		Services:      s.Services,
		FirewallRules: s.FirewallRules,
	}
}
