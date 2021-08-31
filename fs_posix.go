package ruchki

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/user"
	"strconv"
)

type FsPosix struct {
	conf *FsDriverConfig
	dryRun bool
}

func (fs *FsPosix) Setup(c *FsDriverConfig, dryRun bool) error {
	fs.conf = c
	fs.dryRun = dryRun
	return nil
}

func (fs *FsPosix) ApplyDir(d *Directory) error {
	log.WithFields(log.Fields{
		"user": d.User,
		"group": d.Group,
		"permissions": d.Permissions,
	}).Infof("Apply directory `%s`", d.Path)

	if d.Path == "" {
		return fmt.Errorf("directory path is empty")
	}
	if d.User == "" {
		return fmt.Errorf("directory user is empty")
	}
	if d.Group == "" {
		return fmt.Errorf("directory group is empty")
	}
	if d.Permissions == 0 {
		return fmt.Errorf("directory permissions is empty")
	}
	if fs.dryRun {
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
		log.Debugf("Directory `%s` doesn't exists", d.Path)
		err := os.MkdirAll(d.Path, d.Permissions)
		if err != nil  {
			return fmt.Errorf("can't create directory: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("can't check directory: %v", err)
	}

	if err := os.Chmod(d.Path, d.Permissions); err != nil {
		return fmt.Errorf("can't change directory permissions: %v", err)
	}

	if err := os.Chown(d.Path, uid, gid); err != nil {
		return fmt.Errorf("can't change directory permissions: %v", err)
	}

	return nil
}

func (fs *FsPosix) ApplyFile(f *File) error {
	log.WithFields(log.Fields{
		"user": f.User,
		"group": f.Group,
		"permissions": f.Permissions,
	}).Infof("Apply file `%s`", f.Path)
	if f.Path == "" {
		return fmt.Errorf("path is empty")
	}
	if f.FromUri == nil && f.FromTemplate == nil && f.FromContent == nil {
		return fmt.Errorf("uri, template and content are empty")
	}

	if f.User == "" {
		return fmt.Errorf("user is empty")
	}
	if f.Group == "" {
		return fmt.Errorf("group is empty")
	}
	if f.Permissions == 0 {
		return fmt.Errorf("permissions is empty")
	}
	//FromUri url.URL
	//FromTemplate string
	//Vars interface{}
	//FromContent io.ByteReader
	//Compressed bool
	if fs.dryRun {
		return nil
	}

	log.Debugf("Search for user `%s`", f.User)
	us, err := user.Lookup(f.User)
	if err != nil {
		return fmt.Errorf("can't find file user: %v", err)
	}
	uid, err := strconv.Atoi(us.Uid)
	if err != nil {
		return fmt.Errorf("can't parse file user uid: %v", err)
	}

	log.Debugf("Search for group `%s`", f.Group)
	gs, err := user.LookupGroup(f.Group)
	if err != nil {
		return fmt.Errorf("can't find file group: %v", err)
	}
	gid, err := strconv.Atoi(gs.Gid)
	if err != nil {
		return fmt.Errorf("can't parse file group uid: %v", err)
	}

	_, err = os.Stat(f.Path)
	if os.IsNotExist(err) {
		log.Debugf("File `%s` doesn't exists", f.Path)
		err := os.MkdirAll(f.Path, f.Permissions)
		if err != nil  {
			return fmt.Errorf("can't create file: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("can't check file: %v", err)
	}

	if err := os.Chmod(f.Path, f.Permissions); err != nil {
		return fmt.Errorf("can't change file permissions: %v", err)
	}

	if err := os.Chown(f.Path, uid, gid); err != nil {
		return fmt.Errorf("can't change file permissions: %v", err)
	}

	return nil
}
