package posix

import (
	"fmt"
	"github.com/korchasa/kulich/pkg/state"
	"os"

	log "github.com/sirupsen/logrus"
)

func (fs *Posix) AddDir(d *state.Directory) error {
	log.WithFields(log.Fields{
		"user":        d.User,
		"group":       d.Group,
		"permissions": d.Permissions,
	}).Infof("Apply directory `%s`", d.Path)

	uid, gid, err := lookupUsers(d.User, d.Group)
	if err != nil {
		return err
	}

	_, err = os.Stat(d.Path)
	if os.IsNotExist(err) {
		log.Debugf("Create `%s` directory", d.Path)
		if !fs.dryRun {
			err := os.MkdirAll(d.Path, d.Permissions)
			if err != nil {
				return fmt.Errorf("can't create directory: %w", err)
			}
		}
	} else if err != nil {
		return fmt.Errorf("can't check directory: %w", err)
	}

	log.Debugf("Change directory permission to %s", d.Permissions)
	if !fs.dryRun {
		if err := os.Chmod(d.Path, d.Permissions); err != nil {
			return fmt.Errorf("can't change directory permissions: %w", err)
		}
	}

	log.Debugf("Change directory owner to %s(%d):%s:(%d)", d.User, uid, d.Group, gid)
	if !fs.dryRun {
		if err := os.Chown(d.Path, uid, gid); err != nil {
			return fmt.Errorf("can't change directory permissions: %w", err)
		}
	}

	return nil
}
