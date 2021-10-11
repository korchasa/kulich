package yum

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/sysshell"
	log "github.com/sirupsen/logrus"
)

type Yum struct {
	conf *packages.Config
	sh   sysshell.Sysshell
}

func New(conf *packages.Config, sh sysshell.Sysshell) *Yum {
	return &Yum{conf: conf, sh: sh}
}

func (y *Yum) Setup(_ context.Context) error {
	log.Infof("Setup host")
	if y.conf.DryRun {
		return nil
	}
	return nil
}

func (y *Yum) Init(_ context.Context) error {
	log.Infof("Init host")
	if y.conf.DryRun {
		return nil
	}
	res, err := y.sh.Exec(exec.Command("yum", "makecache", "fast"))
	if err != nil {
		return fmt.Errorf("can't update cache: %w", err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("cache update exit with non-zero code `%d`: %s", res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) Add(ctx context.Context, name string) error {
	log.Infof("Install package `%s`", name)
	if y.conf.DryRun {
		return nil
	}

	s := sectionStart("Check package already installed")
	installed, err := y.installed(ctx, name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %w", err)
	}
	sectionEnd(s)
	if installed {
		log.Debugf("Package `%s` already installed", name)
		return nil
	}

	log.Debugf("Run package install")
	res, err := y.sh.Exec(exec.Command("yum", "install", name, "--assumeyes"))
	if err != nil {
		return fmt.Errorf("can't install package `%s`: %w", name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package install exit with non-zero code `%d`: %s", name, res.Exit, res.Stderr)
	}
	return nil
}

type Section struct {
	Name  string
	Start time.Time
}

func sectionStart(n string) *Section {
	log.Debugf("%s...", n)
	return &Section{
		Name:  n,
		Start: time.Now(),
	}
}

func sectionEnd(s *Section) {
	log.Debugf("%s...DONE (%.2fs)", s.Name, time.Since(s.Start).Seconds())
}

func (y *Yum) Remove(ctx context.Context, name string) error {
	log.Infof("Remove package `%s`", name)
	if y.conf.DryRun {
		return nil
	}

	installed, err := y.installed(ctx, name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %w", err)
	}
	if !installed {
		log.Infof("Package `%s` not installed", name)
		return nil
	}

	res, err := y.sh.Exec(exec.Command("yum", "remove", name, "--assumeyes"))
	if err != nil {
		return fmt.Errorf("can't remove package `%s`: %w", name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package remove exit with non-zero code `%d`: %s", name, res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) installed(_ context.Context, name string) (bool, error) {
	res, err := y.sh.Exec(exec.Command("yum", "list", "installed", name, "--assumeyes"))
	if err != nil {
		return false, fmt.Errorf("can't exec yum: %w", err)
	}
	if res.Exit > 1 {
		return false, fmt.Errorf("yum exit with invalid code `%d`: %s", res.Exit, res.Stderr)
	}
	return res.Exit == 0, nil
}
