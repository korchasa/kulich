package posix_test

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/filesystem/posix"
	"github.com/korchasa/kulich/pkg/state"
	"github.com/stretchr/testify/assert"
	"os"
	"time"
)

func (suite *FsIntegrationTestSuite) TestCreateFile_FromLocalFile() {
	t := suite.T()
	src := t.TempDir() + "/src"
	expectedContent := []byte("hello")
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	dst := t.TempDir() + "/dst"
	f := &state.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	}

	err := os.WriteFile(src, expectedContent, 0o600)
	assert.NoError(t, err)

	pfs := new(posix.Posix)
	err = pfs.AddFile(f)

	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, f.Hash)
	}
}

func (suite *FsIntegrationTestSuite) TestCreateFile_ReplaceOldFile() {
	t := suite.T()
	src := fmt.Sprintf("%s/src", t.TempDir())
	dst := fmt.Sprintf("%s/dst", t.TempDir())
	oldContent := []byte("nothing")
	expectedContent := []byte("hello")
	expectedHash := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	f := &state.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	}

	assert.NoError(t, os.WriteFile(dst, oldContent, 0o600))
	assert.NoError(t, os.WriteFile(src, expectedContent, 0o600))
	mt, err := times(dst)
	assert.NoError(t, err)
	time.Sleep(10 * time.Millisecond)

	pfs := new(posix.Posix)
	err = pfs.AddFile(f)

	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, f.Hash)
		nmt, err := times(dst)
		assert.NoError(t, err)
		assert.Greater(t, nmt, mt)
	}
}

func (suite *FsIntegrationTestSuite) TestCreateFile_SuchFileExists() {
	t := suite.T()
	src, dst := fmt.Sprintf("%s/src", t.TempDir()), fmt.Sprintf("%s/dst", t.TempDir())
	expectedContent, expectedHash := []byte("hello"), "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
	nobodyUid, nobodyGid := 99, 99
	f := &state.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o600,
	}

	assert.NoError(t, os.WriteFile(dst, expectedContent, 0o600))
	assert.NoError(t, os.Chown(dst, nobodyUid, nobodyGid))
	assert.NoError(t, os.WriteFile(src, expectedContent, 0o600))
	mt, err := times(dst)
	assert.NoError(t, err)
	time.Sleep(10 * time.Millisecond)

	pfs := new(posix.Posix)
	err = pfs.AddFile(f)

	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, f.Hash)
		nmt, err := times(dst)
		assert.NoError(t, err)
		assert.Equal(t, nmt, mt)
	}
}

func (suite *FsIntegrationTestSuite) TestCreateFile_FromUri() {
	t := suite.T()
	src := "https://github.com/hashicorp/levant/archive/refs/tags/v0.3.0.zip"
	expectedHash := "9d4489776118489c010b49e8001fa93eb94842f99f51f488b44c361a7b007d99"
	dst := t.TempDir() + "/test_from_uri.zip"
	f := &state.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	}

	pfs := new(posix.Posix)
	err := pfs.AddFile(f)

	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, f.Hash)
	}
}

func (suite *FsIntegrationTestSuite) TestCreateFile_FromTemplate() {
	t := suite.T()
	src := t.TempDir() + "/TestLocalTemplate_src.txt"
	srcContent := []byte("hello {{ .name }} with {{ untitle \"sprig\" }}")
	dst := t.TempDir() + "/TestLocalTemplate_dst.txt"
	expectedContent := []byte("hello world with sprig")
	expectedHash := "95a7dff39a9691b61784936f7885610748ede5675fa35f4e2c1487a725108261"
	f := &state.File{
		Path:        dst,
		From:        src,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
		IsTemplate:  true,
		TemplateVars: map[string]string{
			"name": "world",
		},
	}

	err := os.WriteFile(src, srcContent, 0o600)
	assert.NoError(t, err)
	pfs := new(posix.Posix)

	err = pfs.AddFile(f)
	assert.NoError(t, err)
	if assert.FileExists(t, dst) {
		usr, grp, err := getInfo(dst)
		assert.NoError(t, err, "can't get test file info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
		assert.Equal(t, expectedHash, f.Hash)
		actualContent, err := os.ReadFile(dst)
		assert.NoError(t, err)
		assert.Equal(t, expectedContent, actualContent)
	}
}

func times(name string) (mtime float64, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return
	}
	mtime = float64(fi.ModTime().UnixNano()) / 1000000000
	return
}
