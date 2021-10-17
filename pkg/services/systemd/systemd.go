package systemd

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
	"strings"
)

type loadState string
type unitState string
type activeState string
type subState string

const (
	notFound     loadState   = "not-found"
	unitEnabled  unitState   = "enabled"
	unitDisabled unitState   = "disabled"
	active       activeState = "active"
	inactive     activeState = "inactive"
	running      subState    = "running"
	dead         subState    = "dead"
	// failed       subState    = "failed"
)

type Systemd struct {
	dryRun bool
	sh     sysshell.Sysshell
}

func (sys *Systemd) Config(dryRun bool, sh sysshell.Sysshell, opts ...*state.Option) error {
	sys.sh = sh
	sys.dryRun = dryRun
	for _, v := range opts {
		switch v.Name {
		default:
			return fmt.Errorf("unsupported option type `%s`", v.Name)
		}
	}

	return nil
}

func (sys *Systemd) FirstRun() error {
	return nil
}

func (sys *Systemd) BeforeRun() error {
	return nil
}

func (sys *Systemd) AfterRun() error {
	return nil
}

func (sys *Systemd) Add(s *state.Service) error {
	loadState, unitState, activeState, subState, err := sys.serviceState(s.Name)
	if err != nil {
		return fmt.Errorf("can't get service `%s` state: %w", s.Name, err)
	}
	if loadState == notFound {
		return fmt.Errorf("service `%s` doesn't exists", s.Name)
	}
	if s.Disabled {
		if unitState == unitEnabled {
			if err := sys.command("disable", s.Name); err != nil {
				return fmt.Errorf("can't disable service `%s`: %w", s.Name, err)
			}
		}
		if subState == running {
			if err := sys.command("stop", s.Name); err != nil {
				return fmt.Errorf("can't stop service `%s`: %w", s.Name, err)
			}
		}
		return nil
	} else {
		if unitState == unitDisabled {
			if err := sys.command("enable", s.Name); err != nil {
				return fmt.Errorf("can't enable service `%s`: %w", s.Name, err)
			}
		}
		if activeState == inactive {
			if err := sys.command("start", s.Name); err != nil {
				return fmt.Errorf("can't start service `%s`: %w", s.Name, err)
			}
		} else if activeState == active {
			if subState == dead {
				if err := sys.command("restart", s.Name); err != nil {
					return fmt.Errorf("can't restart service `%s`: %w", s.Name, err)
				}
			}
		}
		return nil
	}
}

func (sys *Systemd) Remove(s *state.Service) error {
	loadState, unitState, activeState, _, err := sys.serviceState(s.Name)
	if err != nil {
		return fmt.Errorf("can't get service `%s` state: %w", s.Name, err)
	}
	if loadState == notFound {
		return fmt.Errorf("service `%s` doesn't exists", s.Name)
	}
	if unitState == unitEnabled {
		if err := sys.command("disable", s.Name); err != nil {
			return fmt.Errorf("can't disable service `%s`: %w", s.Name, err)
		}
	}
	if activeState == active {
		if err := sys.command("stop", s.Name); err != nil {
			return fmt.Errorf("can't stop service `%s`: %w", s.Name, err)
		}
	}

	return nil
}

func (sys *Systemd) serviceState(name string) (ls loadState, us unitState, as activeState, ss subState, err error) {
	out, err := sys.sh.SafeExecf("/usr/bin/systemctl show %s.service --no-pager | grep State", name)
	if err != nil {
		err = fmt.Errorf("can't exec systemctl list: %w", err)
		return
	}

	for _, str := range out {
		parts := strings.Split(str, "=")
		if len(parts) != 2 {
			err = fmt.Errorf("unexpected string from systemctl show: %s", str)
			return
		}
		switch parts[0] {
		case "LoadState":
			ls = loadState(parts[1])
			continue
		case "UnitFileState":
			us = unitState(parts[1])
			continue
		case "ActiveState":
			as = activeState(parts[1])
			continue
		case "SubState":
			ss = subState(parts[1])
			continue
		}
	}

	return
}

func (sys *Systemd) command(cmd string, service string) error {
	_, err := sys.sh.SafeExecf("/usr/bin/systemctl %s %s.service", cmd, service)
	if err != nil {
		return fmt.Errorf("can't exec systemctl enable: %w", err)
	}
	return nil
}
