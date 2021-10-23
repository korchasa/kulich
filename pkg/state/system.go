package state

type System struct {
	OsOptions     []Option       `hcl:"os_option,block"`
	Users         []User         `hcl:"user,block"`
	Packages      []Package      `hcl:"package,block"`
	Directories   []Directory    `hcl:"directory,block"`
	Files         []File         `hcl:"file,block"`
	Services      []Service      `hcl:"service,block"`
	FirewallRules []FirewallRule `hcl:"firewall,block"`
}
