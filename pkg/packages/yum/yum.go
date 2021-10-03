package yum

import (
	"context"
	"fmt"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/shell"
	log "github.com/sirupsen/logrus"
)

type Yum struct {
	conf *packages.DriverConfig
	sh   shell.Shell
}

func New(conf *packages.DriverConfig, sh shell.Shell) *Yum {
	return &Yum{conf: conf, sh: sh}
}

func (y *Yum) Setup(_ context.Context) error {
	log.Infof("Setup host")
	if y.conf.DryRun {
		return nil
	}
	return nil
}

func (y *Yum) Init(ctx context.Context) error {
	log.Infof("Init host")
	if y.conf.DryRun {
		return nil
	}
	res, err := y.sh.Exec(ctx, "yum", "makecache", "fast")
	if err != nil {
		return fmt.Errorf("can't update cache: %v", err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("cache update exit with non-zero code `%d`: %s", res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) InstallPackage(ctx context.Context, name string) error {
	log.Infof("Install package `%s`", name)
	if y.conf.DryRun {
		return nil
	}

	installed, err := y.installed(ctx, name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %v", err)
	}
	if installed {
		log.Infof("Package `%s` already installed", name)
		return nil
	}

	res, err := y.sh.Exec(ctx, "yum", "install", name, "--assumeyes")
	if err != nil {
		return fmt.Errorf("can't install package `%s`: %v", name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package install exit with non-zero code `%d`: %s", name, res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) RemovePackage(ctx context.Context, name string) error {
	log.Infof("Install package `%s`", name)
	if y.conf.DryRun {
		return nil
	}

	installed, err := y.installed(ctx, name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %v", err)
	}
	if !installed {
		log.Infof("Package `%s` not installed", name)
		return nil
	}

	res, err := y.sh.Exec(ctx, "yum", "remove", name, "--assumeyes")
	if err != nil {
		return fmt.Errorf("can't remove package `%s`: %v", name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package remove exit with non-zero code `%d`: %s", name, res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) installed(ctx context.Context, name string) (bool, error) {
	res, err := y.sh.Exec(ctx, "yum", "list", "installed", name, "--assumeyes")
	if err != nil {
		return false, fmt.Errorf("can't exec yum: %v", err)
	}
	if res.Exit > 1 {
		return false, fmt.Errorf("yum exit with invalid code `%d`: %s", res.Exit, res.Stderr)
	}
	return res.Exit == 0, nil
}
