package posix

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func getInfo(p string) (*syscall.Stat_t, *user.User, *user.User, error) {
	info, _ := os.Stat(p)
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return nil, nil, nil, fmt.Errorf("can't get test directory `%s` info", p)
	}
	usr, err := user.LookupId(strconv.FormatUint(uint64(stat.Uid), 10))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("can't lookup test directory user: %v", err)
	}
	grp, err := user.LookupId(strconv.FormatUint(uint64(stat.Gid), 10))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("can't lookup test directory group: %v", err)
	}
	return stat, usr, grp, nil
}
