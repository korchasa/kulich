package spec

type Config struct {
	OsDriver         DriverConfig `hcl:"os,block"`
	PackagesDriver   DriverConfig `hcl:"packages,block"`
	FilesystemDriver DriverConfig `hcl:"filesystem,block"`
	ServicesDriver   DriverConfig `hcl:"services,block"`
	FirewallDriver   DriverConfig `hcl:"firewall,block"`
}
