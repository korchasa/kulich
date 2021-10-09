package posix_test

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type FsIntegrationTestSuite struct {
	suite.Suite
}

func (suite *FsIntegrationTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:  true,
		DisableQuote: true,
	})
}

func TestFsIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(FsIntegrationTestSuite))
}

func getInfo(p string) (*user.User, *user.User, error) {
	info, _ := os.Stat(p)
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, nil, fmt.Errorf("can't get test directory `%s` info", p)
	}
	usr, err := user.LookupId(strconv.FormatUint(uint64(stat.Uid), 10))
	if err != nil {
		return nil, nil, fmt.Errorf("can't lookup test directory user: %w", err)
	}
	grp, err := user.LookupId(strconv.FormatUint(uint64(stat.Gid), 10))
	if err != nil {
		return nil, nil, fmt.Errorf("can't lookup test directory group: %w", err)
	}
	return usr, grp, nil
}
