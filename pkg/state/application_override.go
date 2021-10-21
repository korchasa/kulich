package state

type ApplicationOverride struct {
	Name          string                 `hcl:"name,label"`
	OsOptions     []OptionOverride       `hcl:"os_option,block"`
	Users         []UserOverride         `hcl:"user,block"`
	Packages      []PackageOverride      `hcl:"package,block"`
	Directories   []DirectoryOverride    `hcl:"directory,block"`
	Files         []FileOverride         `hcl:"file,block"`
	Services      []ServiceOverride      `hcl:"service,block"`
	FirewallRules []FirewallRuleOverride `hcl:"firewall,block"`
}

func (o *ApplicationOverride) NewApplication() Application {
	return Application{
		Name: o.Name,
		OsOptions: func() (res Options) {
			for _, op := range o.OsOptions {
				res = append(res, op.NewOption())
			}
			return res
		}(),
		Users: func() (res Users) {
			for _, v := range o.Users {
				res = append(res, v.NewUser())
			}
			return res
		}(),
		Packages: func() (res Packages) {
			for _, v := range o.Packages {
				res = append(res, v.NewPackage())
			}
			return res
		}(),
		Directories: func() (res Directories) {
			for _, v := range o.Directories {
				res = append(res, v.NewDirectory())
			}
			return res
		}(),
		Files: func() (res Files) {
			for _, v := range o.Files {
				res = append(res, v.NewFile())
			}
			return res
		}(),
		Services: func() (res Services) {
			for _, v := range o.Services {
				res = append(res, v.NewService())
			}
			return res
		}(),
		FirewallRules: func() (res FirewallRules) {
			for _, v := range o.FirewallRules {
				res = append(res, v.NewRule())
			}
			return res
		}(),
	}
}
