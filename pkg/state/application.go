package state

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

type Applications []Application
