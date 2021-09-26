package posix

import (
	"github.com/korchasa/ruchki/pkg/file_system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CreateDirTestSuite struct {
	suite.Suite
	TestDir string
}

func (suite *CreateDirTestSuite) SetupTest() {
	suite.TestDir = os.Getenv("TEST_DIR")
	if suite.TestDir == "" {
		suite.TestDir = "./tmp"
	}
	_ = os.RemoveAll(suite.TestDir)
}

func (suite *CreateDirTestSuite) TestCreateDir() {
	t := suite.T()
	p := suite.TestDir + "/empty"
	fs := &Posix{}
	err := fs.CreateDir(&file_system.Directory{
		Path:        p,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0755,
	})
	assert.NoError(t, err)
	if assert.DirExists(t, p) {
		_, usr, grp, err := getInfo(p)
		assert.NoError(t, err, "can't get test dir info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
	}
}

func TestCreateDirTestSuite(t *testing.T) {
	suite.Run(t, new(CreateDirTestSuite))
}
