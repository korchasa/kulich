package posix_test

import (
	"github.com/korchasa/ruchki/pkg/filesystem/posix"
	"os"

	"github.com/korchasa/ruchki/pkg/filesystem"
	"github.com/stretchr/testify/assert"
)

func (suite *FsIntegrationTestSuite) TestCopyLocal() {
	t := suite.T()
	src := t.TempDir() + "/local_src.txt"
	expectedContent := []byte("hello")
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	dst := t.TempDir() + "/local_dst.txt"

	err := os.WriteFile(src, expectedContent, 0o600)
	assert.NoError(t, err)
	pfs := posix.NewPosix(&filesystem.DriverConfig{})
	actualHash, err := pfs.CreateFile(&filesystem.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	})
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, actualHash)
	}
}

func (suite *FsIntegrationTestSuite) TestDownload() {
	t := suite.T()
	src := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.0.zip"
	expectedHash := "9d4489776118489c010b49e8001fa93eb94842f99f51f488b44c361a7b007d99"
	dst := t.TempDir() + "/test_from_uri.zip"

	pfs := posix.NewPosix(&filesystem.DriverConfig{})
	actualHash, err := pfs.CreateFile(&filesystem.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	})
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, actualHash)
	}
}

func (suite *FsIntegrationTestSuite) TestLocalTemplate() {
	t := suite.T()
	src := t.TempDir() + "/TestLocalTemplate_src.txt"
	srcContent := []byte("hello {{ .name }} with {{ untitle \"sprig\" }}")
	dst := t.TempDir() + "/TestLocalTemplate_dst.txt"
	expectedContent := []byte("hello world with sprig")
	expectedHash := "95a7dff39a9691b61784936f7885610748ede5675fa35f4e2c1487a725108261"

	err := os.WriteFile(src, srcContent, 0o600)
	assert.NoError(t, err)
	pfs := posix.NewPosix(&filesystem.DriverConfig{})

	actualHash, err := pfs.CreateFile(&filesystem.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
		IsTemplate:  true,
		TemplateVars: map[string]interface{}{
			"name": "world",
		},
	})
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, actualHash)
		actualContent, err := os.ReadFile(dst)
		assert.NoError(t, err)
		assert.Equal(t, expectedContent, actualContent)
	}
}
