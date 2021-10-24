package posix

import (
	"crypto/sha256"
	"fmt"
	"github.com/korchasa/kulich/pkg/spec"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"strconv"
)

func (fs *Posix) AddFile(f *spec.File) error {
	log.Infof("Apply file %s", f.String())
	if fs.dryRun {
		return nil
	}

	log.Debug("Lookup users")
	uid, gid, err := lookupUsers(f.User, f.Group)
	if err != nil {
		return fmt.Errorf("can't lookup user and group: %w", err)
	}

	tmpFilePath, err := fs.prepareTmpFile(f, uid, gid)
	if err != nil {
		return err
	}

	dstExists, err := fileExists(f.Path)
	if err != nil {
		return fmt.Errorf("can't check destination file exists: %w", err)
	}

	log.Debugf("Calculate new file hash")
	newHash, err := calcHash(tmpFilePath)
	if err != nil {
		return fmt.Errorf("can't calculate new file hash: %w", err)
	}
	f.Hash = newHash
	if dstExists {
		oldHash, err := calcHash(f.Path)
		if err != nil {
			return fmt.Errorf("can't calculate old file hash: %w", err)
		}
		if newHash == oldHash {
			log.Info("A similar file already exists")
			_ = os.Remove(tmpFilePath)
			return nil
		}
	}

	if dstExists {
		dstBackupPath := fmt.Sprintf("%s.backup", f.Path)
		_ = os.Remove(dstBackupPath)
		if err := os.Rename(f.Path, dstBackupPath); err != nil {
			return fmt.Errorf("can't backup destination file: %w", err)
		}
	}

	log.Debugf("Rename new version from `%s` to `%s`", tmpFilePath, f.Path)
	if err := os.Rename(tmpFilePath, f.Path); err != nil {
		return fmt.Errorf("can't save file to destination: %w", err)
	}

	return nil
}

func (fs *Posix) prepareTmpFile(f *spec.File, uid int, gid int) (string, error) {
	log.Debugf("Create temp file version for `%s`", f.Path)
	fp, err := createTmpFile(f.Path)
	if err != nil {
		return "", fmt.Errorf("can't create temp file: %w", err)
	}
	defer func() { _ = fp.Close() }()
	log.Debugf("Temp file `%s` created", fp.Name())

	uri, _ := url.Parse(f.From)
	if uri.Scheme != "" {
		log.Debugf("Download content from `%s`", uri.String())
		nb, err := fs.download(fp, uri)
		if err != nil {
			return "", fmt.Errorf("can't download file from url `%s`: %w", uri, err)
		}
		log.Debugf("File downloaded from `%s` (%d bytes)", uri, nb)
	} else {
		log.Debugf("Copy file content from `%s`", f.From)
		nb, err := fs.copy(fp, f.From)
		if err != nil {
			return "", fmt.Errorf("can't copy file from `%s`: %w", f.From, err)
		}
		log.Debugf("File copied from `%s` (%d bytes)", f.From, nb)
	}

	if f.IsTemplate {
		log.Debugf("Render template")
		err := fs.render(fp, f.TemplateVars)
		if err != nil {
			return "", fmt.Errorf("can't render template: %w", err)
		}
	}

	log.Debugf("Change file permissions to %s", f.Permissions)
	if err := fp.Chmod(f.Permissions); err != nil {
		return "", fmt.Errorf("can't change file permissions: %w", err)
	}

	log.Debugf("Change directory owner to %s(%d):%s:(%d)", f.User, uid, f.Group, gid)
	if err := fp.Chown(uid, gid); err != nil {
		return "", fmt.Errorf("can't change file permissions: %w", err)
	}

	if err := fp.Close(); err != nil {
		return "", fmt.Errorf("can't close file: %w", err)
	}

	return fp.Name(), nil
}

func lookupUsers(usr, grp string) (int, int, error) {
	log.Debugf("Search for user `%s`", usr)
	us, err := user.Lookup(usr)
	if err != nil {
		return 0, 0, fmt.Errorf("can't find file user: %w", err)
	}
	uid, err := strconv.Atoi(us.Uid)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse file user uid: %w", err)
	}

	log.Debugf("Search for group `%s`", grp)
	gs, err := user.LookupGroup(grp)
	if err != nil {
		return 0, 0, fmt.Errorf("can't find file group: %w", err)
	}
	gid, err := strconv.Atoi(gs.Gid)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse file group uid: %w", err)
	}

	return uid, gid, nil
}

func createTmpFile(path string) (*os.File, error) {
	tmpName := fmt.Sprintf("%s.tmp", path)
	fileExists, err := fileExists(tmpName)
	if err != nil {
		return nil, fmt.Errorf("can't check temp file already exists: %w", err)
	}
	if fileExists {
		log.Debugf("Temp file `%s` already exists", tmpName)
		if err := os.Remove(tmpName); err != nil {
			return nil, fmt.Errorf("can't delete temp file `%s`: %w", tmpName, err)
		}
	}

	f, err := os.Create(tmpName)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %w", err)
	}
	return f, nil
}

func calcHash(path string) (string, error) {
	h := sha256.New()
	s, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("can't read file content: %w", err)
	}
	h.Write(s)
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
