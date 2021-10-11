package packages

import (
	"context"
)

type Packages interface {
	Setup(ctx context.Context, c *Config) error
	Init(ctx context.Context, c *Config) error
	Add(ctx context.Context, name string) error
	Remove(ctx context.Context, name string) error
}

type Config struct {
	DryRun bool
}
