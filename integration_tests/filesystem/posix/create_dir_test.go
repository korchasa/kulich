package posix_test

import (
	"github.com/korchasa/kulich/pkg/filesystem/posix"
	"github.com/korchasa/kulich/pkg/spec"
	"github.com/stretchr/testify/assert"
)

func (suite *FsIntegrationTestSuite) TestCreateDir() {
	t := suite.T()
	p := t.TempDir() + "/empty"
	pfs := new(posix.Posix)
	err := pfs.AddDir(&spec.Directory{
		Path:        p,
		User:        "nobody",
		Group:       "nobody",
		Permissions: 0o755,
	})
	assert.NoError(t, err)
	if assert.DirExists(t, p) {
		usr, grp, err := getInfo(p)
		assert.NoError(t, err, "can't get test dir info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
	}
}
