package state

type State struct {
	Type Type `hcl:"type,block"`
}

type Type struct {
	Name         string        `hcl:"name,label"`
	Config       ServerConfig  `hcl:"config,block"`
	System       System        `hcl:"system,block"`
	Applications []Application `hcl:"application,block"`
	Servers      []Server      `hcl:"server,block"`
}

type Server struct {
	Name         string                `hcl:"name,label"`
	System       *SystemOverride       `hcl:"system,block"`
	Applications []ApplicationOverride `hcl:"application,block"`
}

type ServerConfig struct {
	OsDriver         DriverConfig `hcl:"os,block"`
	PackagesDriver   DriverConfig `hcl:"packages,block"`
	FilesystemDriver DriverConfig `hcl:"filesystem,block"`
	ServicesDriver   DriverConfig `hcl:"services,block"`
	FirewallDriver   DriverConfig `hcl:"firewall,block"`
}

type DriverConfig struct {
	Name string `hcl:"name,label"`
}

type System struct {
	OsOptions     []Option    `hcl:"os_option,block"`
	Users         []User      `hcl:"user,block"`
	Packages      []Package   `hcl:"package,block"`
	Directories   []Directory `hcl:"directory,block"`
	Files         []File      `hcl:"file,block"`
	Services      []Service   `hcl:"service,block"`
	FirewallRules []Rule      `hcl:"firewall,block"`
}

type SystemOverride struct {
	OsOptions     []Option    `hcl:"os_option,block"`
	Users         []User      `hcl:"user,block"`
	Packages      []Package   `hcl:"package,block"`
	Directories   []Directory `hcl:"directory,block"`
	Files         []File      `hcl:"file,block"`
	Services      []Service   `hcl:"service,block"`
	FirewallRules []Rule      `hcl:"firewall,block"`
}

type Application struct {
	Name          string              `hcl:"name,label"`
	OsOptions     []OptionOverride    `hcl:"os_option,block"`
	Users         []UserOverride      `hcl:"user,block"`
	Packages      []PackageOverride   `hcl:"package,block"`
	Directories   []DirectoryOverride `hcl:"directory,block"`
	Files         []FileOverride      `hcl:"file,block"`
	Services      []ServiceOverride   `hcl:"service,block"`
	FirewallRules []RuleOverride      `hcl:"firewall,block"`
}

type ApplicationOverride struct {
	Name          string              `hcl:"name,label"`
	OsOptions     []OptionOverride    `hcl:"os_option,block"`
	Users         []UserOverride      `hcl:"user,block"`
	Packages      []PackageOverride   `hcl:"package,block"`
	Directories   []DirectoryOverride `hcl:"directory,block"`
	Files         []FileOverride      `hcl:"file,block"`
	Services      []ServiceOverride   `hcl:"service,block"`
	FirewallRules []RuleOverride      `hcl:"firewall,block"`
}
