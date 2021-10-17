package centos7_test

import (
	"github.com/korchasa/ruchki/pkg/os"
	"github.com/korchasa/ruchki/pkg/os/centos7"
	"github.com/korchasa/ruchki/pkg/sysshell/posix"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"

	log "github.com/sirupsen/logrus"
)

type Centos7IntegrationTestSuite struct {
	suite.Suite
}

func (suite *Centos7IntegrationTestSuite) SetupTest() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestIptablesIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(Centos7IntegrationTestSuite))
}

func (suite *Centos7IntegrationTestSuite) TestAddRemove() {
	sh := posix.New()
	osm := new(centos7.Centos7)
	assert.NoError(suite.T(), osm.Config(false, sh))

	u := &os.User{Name: "alice", System: false}
	assert.NoError(suite.T(), osm.AddUser(u))
	assert.DirExists(suite.T(), "/home/alice/")

	assert.NoError(suite.T(), osm.RemoveUser(u))
	assert.NoDirExists(suite.T(), "/home/alice")
}

func (suite *Centos7IntegrationTestSuite) TestAddRemoveSystem() {
	sh := posix.New()
	osm := new(centos7.Centos7)
	assert.NoError(suite.T(), osm.Config(false, sh))

	systemUser := &os.User{Name: "bob", System: true}
	assert.NoError(suite.T(), osm.AddUser(systemUser))
	assert.NoDirExists(suite.T(), "/home/bob")

	assert.NoError(suite.T(), osm.RemoveUser(systemUser))
	assert.NoDirExists(suite.T(), "/home/bob")
}
