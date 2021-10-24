package yum_test

import (
	"github.com/korchasa/kulich/pkg/packages/yum"
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/korchasa/kulich/pkg/sysshell/posix"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// var execMock func(path string, args ...string) (*shell.Result, error)
//
// type shellMock struct{}
//
// func (u shellMock) Exec(_ context.Context, path string, args ...string) (*shell.Result, error) {
//	 return execMock(path, args...)
// }

type YumIntegrationTestSuite struct {
	suite.Suite
}

func (suite *YumIntegrationTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestYumIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(YumIntegrationTestSuite))
}

func (suite *YumIntegrationTestSuite) TestYum_BeforeRun() {
	sh := posix.New()
	mng := new(yum.Yum)
	assert.NoError(suite.T(), mng.Config(false, sh))
	err := mng.BeforeRun()
	assert.NoError(suite.T(), err)
}

func (suite *YumIntegrationTestSuite) TestYum_RemovePackage() {
	sh := posix.New()
	mng := new(yum.Yum)
	assert.NoError(suite.T(), mng.Config(false, sh))
	assert.NoError(suite.T(), mng.Remove(&spec.Package{Name: "epel-release"}))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage() {
	sh := posix.New()
	mng := new(yum.Yum)
	assert.NoError(suite.T(), mng.Config(false, sh))
	assert.NoError(suite.T(), mng.Add(&spec.Package{Name: "epel-release"}))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage_Repeat() {
	sh := posix.New()
	mng := new(yum.Yum)
	assert.NoError(suite.T(), mng.Config(false, sh))
	assert.NoError(suite.T(), mng.Add(&spec.Package{Name: "epel-release"}))
	assert.NoError(suite.T(), mng.Add(&spec.Package{Name: "epel-release"}))
}
