package packages

import (
	"context"
)

type Driver interface {
	Setup(ctx context.Context, c *DriverConfig) error
	Init(ctx context.Context, c *DriverConfig) error
	InstallPackage(ctx context.Context, name string) error
	RemovePackage(ctx context.Context, name string) error
}

type DriverConfig struct {
	DryRun bool
}
