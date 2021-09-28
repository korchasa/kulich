package posix

import (
	"fmt"
	"github.com/korchasa/ruchki/pkg/fs"
	log "github.com/sirupsen/logrus"
	"os"
	"os/user"
	"strconv"
)

func (fs *Posix) CreateDir(d *fs.Directory) error {
	log.WithFields(log.Fields{
		"user":        d.User,
		"group":       d.Group,
		"permissions": d.Permissions,
	}).Infof("Apply directory `%s`", d.Path)

	if fs.conf.DryRun {
		return nil
	}

	log.Debugf("Search for user `%s`", d.User)
	dus, err := user.Lookup(d.User)
	if err != nil {
		return fmt.Errorf("can't find directory user: %v", err)
	}
	uid, err := strconv.Atoi(dus.Uid)
	if err != nil {
		return fmt.Errorf("can't parse directory user uid: %v", err)
	}

	log.Debugf("Search for group `%s`", d.Group)
	dgs, err := user.LookupGroup(d.Group)
	if err != nil {
		return fmt.Errorf("can't find directory group: %v", err)
	}
	gid, err := strconv.Atoi(dgs.Gid)
	if err != nil {
		return fmt.Errorf("can't parse directory group uid: %v", err)
	}

	_, err = os.Stat(d.Path)
	if os.IsNotExist(err) {
		log.Debugf("Create `%s` directory", d.Path)
		err := os.MkdirAll(d.Path, d.Permissions)
		if err != nil {
			return fmt.Errorf("can't create directory: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("can't check directory: %v", err)
	}

	log.Debugf("Change directory permission to %s", d.Permissions)
	if err := os.Chmod(d.Path, d.Permissions); err != nil {
		return fmt.Errorf("can't change directory permissions: %v", err)
	}

	log.Debugf("Change directory owner to %s(%d):%s:(%d)", d.User, uid, d.Group, gid)
	if err := os.Chown(d.Path, uid, gid); err != nil {
		return fmt.Errorf("can't change directory permissions: %v", err)
	}

	return nil
}
