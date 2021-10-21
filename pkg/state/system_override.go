package state

type SystemOverride struct {
	OsOptions     []OptionOverride       `hcl:"os_option,block"`
	Users         []UserOverride         `hcl:"user,block"`
	Packages      []PackageOverride      `hcl:"package,block"`
	Directories   []DirectoryOverride    `hcl:"directory,block"`
	Files         []FileOverride         `hcl:"file,block"`
	Services      []ServiceOverride      `hcl:"service,block"`
	FirewallRules []FirewallRuleOverride `hcl:"firewall,block"`
}
