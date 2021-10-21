package state

type Application struct {
	Name          string        `hcl:"name,label"`
	OsOptions     Options       `hcl:"os_option,block"`
	Users         Users         `hcl:"user,block"`
	Packages      Packages      `hcl:"package,block"`
	Directories   Directories   `hcl:"directory,block"`
	Files         Files         `hcl:"file,block"`
	Services      Services      `hcl:"service,block"`
	FirewallRules FirewallRules `hcl:"firewall,block"`
}

func (app *Application) Apply(override ApplicationOverride) bool {
	if app.Name != override.Name {
		return false
	}
	app.OsOptions.Apply(override.OsOptions)
	app.Users.Apply(override.Users)
	app.Packages.Apply(override.Packages)
	app.Directories.Apply(override.Directories)
	app.Files.Apply(override.Files)
	app.Services.Apply(override.Services)
	app.FirewallRules.Apply(override.FirewallRules)
	return true
}
