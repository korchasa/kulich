package packages

import (
	"context"
)

type Driver interface {
	Setup(ctx context.Context, c *DriverConfig) error
	Init(ctx context.Context, c *DriverConfig) error
	Add(ctx context.Context, name string) error
	Remove(ctx context.Context, name string) error
}

type DriverConfig struct {
	DryRun bool
}
