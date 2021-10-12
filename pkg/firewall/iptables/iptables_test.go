package iptables_test

import (
	"github.com/korchasa/ruchki/pkg/firewall"
	"github.com/korchasa/ruchki/pkg/firewall/iptables"
	"github.com/korchasa/ruchki/pkg/sysshell/posix"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"

	log "github.com/sirupsen/logrus"
)

type IptablesIntegrationTestSuite struct {
	suite.Suite
}

func (suite *IptablesIntegrationTestSuite) SetupTest() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestIptablesIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IptablesIntegrationTestSuite))
}

func (suite *IptablesIntegrationTestSuite) TestAdd() {
	sh := posix.New()
	ipt := iptables.New(&firewall.Config{}, sh)
	err := ipt.Add(&firewall.Rule{Ports: []string{"10001", "10100:10200"}, Targets: []string{"192.158.1.1", "192.158.1.0/24"}})
	assert.NoError(suite.T(), err)
}

func (suite *IptablesIntegrationTestSuite) TestRemove() {
	sh := posix.New()
	ipt := iptables.New(&firewall.Config{}, sh)
	err := ipt.Add(&firewall.Rule{Ports: []string{"20000", "20100:20200"}, Targets: []string{"192.158.2.1", "192.158.2.0/24"}})
	assert.NoError(suite.T(), err)
	err = ipt.Remove(&firewall.Rule{Ports: []string{"20000"}, Targets: []string{"192.158.2.1"}})
	assert.NoError(suite.T(), err)
	err = ipt.Remove(&firewall.Rule{Ports: []string{"20100:20200"}, Targets: []string{"192.158.2.0/24"}})
	assert.NoError(suite.T(), err)
}
