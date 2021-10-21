package state

type System struct {
	OsOptions     Options       `hcl:"os_option,block"`
	Users         Users         `hcl:"user,block"`
	Packages      Packages      `hcl:"package,block"`
	Directories   Directories   `hcl:"directory,block"`
	Files         Files         `hcl:"file,block"`
	Services      Services      `hcl:"service,block"`
	FirewallRules FirewallRules `hcl:"firewall,block"`
}

func (sys *System) Apply(override *SystemOverride) {
	sys.OsOptions.Apply(override.OsOptions)
	sys.Users.Apply(override.Users)
	sys.Packages.Apply(override.Packages)
	sys.Directories.Apply(override.Directories)
	sys.Files.Apply(override.Files)
	sys.Services.Apply(override.Services)
	sys.FirewallRules.Apply(override.FirewallRules)
}
