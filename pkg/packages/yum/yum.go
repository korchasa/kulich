package yum

import (
	"context"
	"fmt"
	"github.com/korchasa/kulich/pkg/spec"
	"os/exec"
	"time"

	"github.com/korchasa/kulich/pkg/sysshell"
	log "github.com/sirupsen/logrus"
)

type Yum struct {
	sh     sysshell.Sysshell
	dryRun bool
}

func (y *Yum) Config(dryRun bool, sh sysshell.Sysshell, opts ...*spec.OsOption) error {
	y.sh = sh
	y.dryRun = dryRun
	for _, v := range opts {
		switch v.Name {
		default:
			return fmt.Errorf("unsupported option type `%s`", v.Name)
		}
	}

	return nil
}

func (y *Yum) FirstRun() error {
	return nil
}

func (y *Yum) BeforeRun() error {
	log.Infof("Yum init")
	if y.dryRun {
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

func (y *Yum) AfterRun() error {
	return nil
}

func (y *Yum) Add(p *spec.Package) error {
	log.Infof("Install package `%s`", p.Name)
	if y.dryRun {
		return nil
	}

	s := sectionStart("Check package already installed")
	installed, err := y.installed(context.TODO(), p.Name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %w", err)
	}
	sectionEnd(s)
	if installed {
		log.Debugf("Package `%s` already installed", p.Name)
		return nil
	}

	log.Debugf("Run package install")
	res, err := y.sh.Exec(exec.Command("yum", "install", p.Name, "--assumeyes"))
	if err != nil {
		return fmt.Errorf("can't install package `%s`: %w", p.Name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package install exit with non-zero code `%d`: %s", p.Name, res.Exit, res.Stderr)
	}
	return nil
}

func (y *Yum) Remove(p *spec.Package) error {
	log.Infof("Removed package `%s`", p.Name)
	if y.dryRun {
		return nil
	}

	installed, err := y.installed(context.TODO(), p.Name)
	if err != nil {
		return fmt.Errorf("can't check package installed or not: %w", err)
	}
	if !installed {
		log.Infof("Package `%s` not installed", p.Name)
		return nil
	}

	res, err := y.sh.Exec(exec.Command("yum", "remove", p.Name, "--assumeyes"))
	if err != nil {
		return fmt.Errorf("can't remove package `%s`: %w", p.Name, err)
	}
	if res.Exit != 0 {
		return fmt.Errorf("`%s` package remove exit with non-zero code `%d`: %s", p.Name, res.Exit, res.Stderr)
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
