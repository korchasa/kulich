package spec

type Config struct {
	OsDriver         DriverConfig
	PackagesDriver   DriverConfig
	FilesystemDriver DriverConfig
	ServicesDriver   DriverConfig
	FirewallDriver   DriverConfig
}
