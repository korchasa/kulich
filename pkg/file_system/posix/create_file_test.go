package posix

import (
	"github.com/korchasa/ruchki/pkg/file_system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type CreateFileTestSuite struct {
	suite.Suite
	TestDir string
}

func (suite *CreateFileTestSuite) SetupTest() {
	suite.TestDir = os.Getenv("TEST_DIR")
	if suite.TestDir == "" {
		suite.TestDir = "./tmp"
	}
	_ = os.RemoveAll(suite.TestDir)
	_ = os.MkdirAll(suite.TestDir, 0755)
}

func (suite *CreateFileTestSuite) TestCopyLocal() {
	t := suite.T()
	src := suite.TestDir + "/local_src.txt"
	expectedContent := []byte("hello")
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	dst := suite.TestDir + "/local_dst.txt"

	err := os.WriteFile(src, expectedContent, 0644)
	assert.NoError(t, err)
	pfs := &Posix{}
	actualHash, err := pfs.CreateFile(&file_system.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0755,
	})
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		_, usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, actualHash)
	}
}

func (suite *CreateFileTestSuite) TestDownload() {
	t := suite.T()
	src := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.0.zip"
	expectedHash := "9d4489776118489c010b49e8001fa93eb94842f99f51f488b44c361a7b007d99"
	dst := suite.TestDir + "/test_from_uri.zip"

	pfs := &Posix{}
	actualHash, err := pfs.CreateFile(&file_system.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0755,
	})
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		_, usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, actualHash)
	}
}

func TestCreateFileTestSuite(t *testing.T) {
	suite.Run(t, new(CreateFileTestSuite))
}
