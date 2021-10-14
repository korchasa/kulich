package iptables_test

import (
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/firewall/iptables"
	"github.com/korchasa/ruchki/pkg/sysshell"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IptablesTestSuite struct {
	suite.Suite
}

func (suite *IptablesTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestSystemdIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IptablesTestSuite))
}

func (suite *IptablesTestSuite) TestImplementInterface() {
	var _ firewall.Firewall = (*iptables.Iptables)(nil)
}

func (suite *IptablesTestSuite) TestSystemd_Add() {
	sh := new(sysshell.Mock)
	ipt := new(iptables.Iptables)
	assert.NoError(suite.T(), ipt.Config(false, sh))

	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 --src 192.168.0.1 -m comment --comment "2ff487a2" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 --src 192.168.100.1/24 -m comment --comment "9e4244a6" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 --src 192.168.0.1 -m comment --comment "c3fb77c9" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 --src 192.168.100.1/24 -m comment --comment "c5504be1" -j ACCEPT`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 2222 -m comment --comment "a9cdce76" -j DROP`)
	expectSafeExec(sh, `iptables --append INPUT --protocol tcp --dport 1000:2000 -m comment --comment "4d6dc86e" -j DROP`)

	err := ipt.Add(&firewall.Rule{
		Ports:   []string{"2222", "1000:2000"},
		Targets: []string{"192.168.0.1", "192.168.100.1/24"},
	})
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}

func (suite *IptablesTestSuite) TestSystemd_Remove() {
	sh := new(sysshell.Mock)
	ipt := new(iptables.Iptables)
	assert.NoError(suite.T(), ipt.Config(false, sh))

	expectSafeExec(sh, `iptables --delete INPUT --protocol tcp --dport 2222 --src 192.168.100.1/24 -m comment --comment "9e4244a6" -j ACCEPT`)
	expectSafeExec(sh, `iptables --delete INPUT --protocol tcp --dport 2222 -m comment --comment "a9cdce76" -j DROP`)

	err := ipt.Remove(&firewall.Rule{
		Ports:   []string{"2222"},
		Targets: []string{"192.168.100.1/24"},
	})
	assert.NoError(suite.T(), err)
	sh.AssertExpectationsInOrder(suite.T())
}

func expectSafeExec(sh *sysshell.Mock, text string) {
	sh.On("SafeExec", text).Return([]string{}, nil).Once()
}
