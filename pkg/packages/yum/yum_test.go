package yum_test

import (
	"context"
	"github.com/korchasa/ruchki/pkg/packages/yum"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/sysshell/posix"
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

func (suite *YumIntegrationTestSuite) TestYum_Init() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	err := mng.Init(context.Background())
	assert.NoError(suite.T(), err)
}

func (suite *YumIntegrationTestSuite) TestYum_RemovePackage() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.RemovePackage(context.Background(), "epel-release"))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage_Repeat() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
}
