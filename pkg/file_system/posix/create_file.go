package posix

import (
	"crypto/sha256"
	"fmt"
	"github.com/korchasa/ruchki/pkg/file_system"
	log "github.com/sirupsen/logrus"
	"io"
	"net/url"
	"os"
	"os/user"
	"strconv"
)

func (fs *Posix) CreateFile(s *file_system.File) (string, error) {
	log.Infof("Apply file %s", s.String())
	if fs.dryRun {
		return "", nil
	}

	uid, gid, err := lookupUsers(s)
	if err != nil {
		return "", err
	}

	fp, err := ensureFile(s.Path)
	if err != nil {
		return "", fmt.Errorf("can't ensure file: %v", err)
	}
	defer fp.Close()

	uri, _ := url.Parse(s.From)
	if uri.Scheme != "" {
		nb, path, err := Download(uri, fp)
		if err != nil {
			return "", fmt.Errorf("can't download file from url `%s`: %v", uri, err)
		}
		log.Debugf("File downloaded from `%s` to `%s` (%d bytes)", uri, s.Path, nb)
		s.From = path
	} else {
		nb, err := Copy(s.From, fp)
		if err != nil {
			return "", fmt.Errorf("can't copy file from `%s`: %v", s.From, err)
		}
		log.Debugf("File copied from `%s` to `%s` (%d bytes)", s.From, s.Path, nb)
	}

	if err := fp.Chmod(s.Permissions); err != nil {
		return "", fmt.Errorf("can't change file permissions: %v", err)
	}

	if err := fp.Chown(uid, gid); err != nil {
		return "", fmt.Errorf("can't change file permissions: %v", err)
	}

	hash, err := calcHash(fp)
	if err != nil {
		return "", fmt.Errorf("can't calculate sha256: %v", err)
	}

	if err := fp.Close(); err != nil {
		return "", fmt.Errorf("can't close file: %v", err)
	}

	return hash, nil
}

func lookupUsers(s *file_system.File) (int, int, error) {
	log.Debugf("Search for user `%s`", s.User)
	us, err := user.Lookup(s.User)
	if err != nil {
		return 0, 0, fmt.Errorf("can't find file user: %v", err)
	}
	uid, err := strconv.Atoi(us.Uid)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse file user uid: %v", err)
	}

	log.Debugf("Search for group `%s`", s.Group)
	gs, err := user.LookupGroup(s.Group)
	if err != nil {
		return 0, 0, fmt.Errorf("can't find file group: %v", err)
	}
	gid, err := strconv.Atoi(gs.Gid)
	if err != nil {
		return 0, 0, fmt.Errorf("can't parse file group uid: %v", err)
	}
	return uid, gid, nil
}

func ensureFile(path string) (*os.File, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Create(path)
	}
	if err != nil {
		return nil, fmt.Errorf("can't get file stat: %v", err)
	}
	if info.IsDir() {
		if err := os.Remove(path); err != nil {
			return nil, fmt.Errorf("can't remove directory: %v", err)
		}
	}
	return os.Open(path)
}

func calcHash(f *os.File) (string, error) {
	if _, err := f.Seek(0, 0); err != nil {
		return "", fmt.Errorf("can't return to begin of file: %v", err)
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("can't read file content: %v", err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}