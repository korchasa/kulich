package posix

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"strconv"

	"github.com/korchasa/ruchki/pkg/filesystem"
	log "github.com/sirupsen/logrus"
)

func (fs *Posix) CreateFile(s *filesystem.File) (string, error) {
	log.Infof("Apply file %s", s.String())
	if fs.conf.DryRun {
		return "", nil
	}

	uid, gid, err := lookupUsers(s.User, s.Group)
	if err != nil {
		return "", err
	}

	tmpFilePath, err := fs.prepareTmpFile(s, uid, gid)
	if err != nil {
		return "", err
	}

	log.Debugf("Calculate file hash")
	hash, err := calcHash(tmpFilePath)
	if err != nil {
		return "", fmt.Errorf("can't calculate sha256: %w", err)
	}
	if s.Sha256Hash != "" {
		if hash == s.Sha256Hash {
			log.Debugf("New file has the same hash as destination")
		}
	}

	dstExists, err := fileExists(s.Path)
	if err != nil {
		return "", fmt.Errorf("can't check destination file exists: %w", err)
	}
	if dstExists {
		dstBackupPath := fmt.Sprintf("%s.backup", s.Path)
		_ = os.Remove(dstBackupPath)
		if err := os.Rename(s.Path, dstBackupPath); err != nil {
			return "", fmt.Errorf("can't backup destination file: %w", err)
		}
	}

	log.Debugf("Rename new version from `%s` to `%s`", tmpFilePath, s.Path)
	if err := os.Rename(tmpFilePath, s.Path); err != nil {
		return "", fmt.Errorf("can't save file to destination: %w", err)
	}

	return hash, nil
}

func (fs *Posix) prepareTmpFile(s *filesystem.File, uid int, gid int) (string, error) {
	log.Debugf("Create temp file version for `%s`", s.Path)
	fp, err := createTmpFile(s.Path)
	if err != nil {
		return "", fmt.Errorf("can't create temp file: %w", err)
	}
	defer fp.Close()
	log.Debugf("Temp file `%s` created", fp.Name())

	uri, _ := url.Parse(s.From)
	if uri.Scheme != "" {
		log.Debugf("Download content from `%s`", uri.String())
		nb, err := fs.download(fp, uri)
		if err != nil {
			return "", fmt.Errorf("can't download file from url `%s`: %w", uri, err)
		}
		log.Debugf("File downloaded from `%s` (%d bytes)", uri, nb)
	} else {
		log.Debugf("Copy file content from `%s`", s.From)
		nb, err := fs.copy(fp, s.From)
		if err != nil {
			return "", fmt.Errorf("can't copy file from `%s`: %w", s.From, err)
		}
		log.Debugf("File copied from `%s` (%d bytes)", s.From, nb)
	}

	if s.IsTemplate {
		log.Debugf("Render template")
		err := fs.render(fp, s.TemplateVars)
		if err != nil {
			return "", fmt.Errorf("can't render template: %w", err)
		}
	}

	log.Debugf("Change file permissions to %s", s.Permissions)
	if err := fp.Chmod(s.Permissions); err != nil {
		return "", fmt.Errorf("can't change file permissions: %w", err)
	}

	log.Debugf("Change directory owner to %s(%d):%s:(%d)", s.User, uid, s.Group, gid)
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
