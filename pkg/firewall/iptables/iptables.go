package iptables

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/sysshell"
	"net"
	"strconv"
	"strings"
)

type Iptables struct {
	conf *firewall.Config
	sh   sysshell.Sysshell
}

func New(conf *firewall.Config, sh sysshell.Sysshell) *Iptables {
	return &Iptables{conf: conf, sh: sh}
}

func (i *Iptables) FirstRun() error {
	return nil
}

func (i *Iptables) BeforeRun() error {
	return nil
}

func (i *Iptables) AfterRun() error {
	return nil
}

func (i *Iptables) Add(r *firewall.Rule) error {
	return i.cmd(r, "append")
}

func (i *Iptables) Remove(r *firewall.Rule) error {
	return i.cmd(r, "delete")
}

func (i *Iptables) cmd(r *firewall.Rule, cmd string) error {
	protocol := r.Protocol
	if protocol == "" {
		protocol = firewall.DefaultProtocol
	}

	for _, port := range r.Ports {
		if !validPort(port) {
			return fmt.Errorf("can't parse port or ports range `%s`", port)
		}
		for _, target := range r.Targets {
			if !validTarget(target) {
				return fmt.Errorf("can't parse target `%s`", target)
			}
			if !r.IsOutput {
				if _, err := i.sh.SafeExecf(
					"iptables --%s INPUT --protocol %s --dport %s --src %s -m comment --comment \"%s\" -j ACCEPT",
					cmd,
					protocol,
					port,
					target,
					identifier(r, target, port),
				); err != nil {
					return fmt.Errorf("can't %s accept target input rule: %w", cmd, err)
				}
			} else {
				if _, err := i.sh.SafeExecf(
					"iptables --%s OUTPUT --protocol %s --dport %s --src %s -m comment --comment \"%s\" -j ACCEPT",
					cmd,
					protocol,
					port,
					target,
					identifier(r, target, port),
				); err != nil {
					return fmt.Errorf("can't %s accept target output rule: %w", cmd, err)
				}
			}
		}
	}
	for _, port := range r.Ports {
		if !validPort(port) {
			return fmt.Errorf("can't parse port or ports range `%s`", port)
		}
		if !r.IsOutput {
			if _, err := i.sh.SafeExecf(
				"iptables --%s INPUT --protocol %s --dport %s -m comment --comment \"%s\" -j DROP",
				cmd,
				protocol,
				port,
				identifier(r, "", port),
			); err != nil {
				return fmt.Errorf("can't %s drop input rule: %w", cmd, err)
			}
		} else {
			if _, err := i.sh.SafeExecf(
				"iptables --%s OUTPUT --protocol %s --dport %s -m comment --comment \"%s\" -j DROP",
				cmd,
				protocol,
				port,
				identifier(r, "", port),
			); err != nil {
				return fmt.Errorf("can't %s drop output rule: %w", cmd, err)
			}
		}
	}
	return nil
}

func validTarget(target string) bool {
	if net.ParseIP(target) != nil {
		return true
	}
	_, n, _ := net.ParseCIDR(target)
	if n != nil {
		return true
	}

	return false
}

func validPort(port string) bool {
	for _, p := range strings.Split(port, ":") {
		_, err := strconv.Atoi(p)
		if err != nil {
			return false
		}
	}

	return true
}

func identifier(r *firewall.Rule, target, port string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s-%s-%v-%s-%s", r.Identifier, r.Protocol, r.IsOutput, target, port)))
	return hex.EncodeToString(hash[:])[0:8]
}
