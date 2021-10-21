package iptables_test

import (
	"github.com/korchasa/kulich/pkg/firewall"
	"github.com/korchasa/kulich/pkg/firewall/iptables"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/korchasa/kulich/pkg/sysshell"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImplementInterface(t *testing.T) {
	var _ firewall.Firewall = (*iptables.Iptables)(nil)
}

func TestSystemd_Add(t *testing.T) {
	sh := new(sysshell.Mock)
	ipt := new(iptables.Iptables)
	assert.NoError(t, ipt.Config(false, sh))

	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 --src 192.168.0.1 -m comment --comment "2ff487a2" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 --src 192.168.100.1/24 -m comment --comment "9e4244a6" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 --src 192.168.0.1 -m comment --comment "c3fb77c9" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 --src 192.168.100.1/24 -m comment --comment "c5504be1" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 -m comment --comment "a9cdce76" -j DROP`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 -m comment --comment "4d6dc86e" -j DROP`)

	err := ipt.Add(&state.FirewallRule{
		Ports:   []string{"2222", "1000:2000"},
		Targets: []string{"192.168.0.1", "192.168.100.1/24"},
	})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}

func TestSystemd_Remove(t *testing.T) {
	sh := new(sysshell.Mock)
	ipt := new(iptables.Iptables)
	assert.NoError(t, ipt.Config(false, sh))

	expectSafeExec(sh, `iptables --delete INPUT --protocol tcp --dport 2222 --src 192.168.100.1/24 -m comment --comment "9e4244a6" -j ACCEPT`)
	expectSafeExec(sh, `iptables --delete INPUT --protocol tcp --dport 2222 -m comment --comment "a9cdce76" -j DROP`)

	err := ipt.Remove(&state.FirewallRule{
		Ports:   []string{"2222"},
		Targets: []string{"192.168.100.1/24"},
	})
	assert.NoError(t, err)
	sh.AssertExpectationsInOrder(t)
}

func expectSafeExec(sh *sysshell.Mock, text string) {
	sh.On("SafeExec", text).Return([]string{}, nil).Once()
}
