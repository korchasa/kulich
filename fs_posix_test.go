package ruchki

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"testing"
)

func TestFsLinux_ApplyDir_Errors(t *testing.T) {
	tests := []struct {
		errMsg string
		conf *Directory
	}{
		{"directory path is empty", &Directory{"", "nobody", "nobody", 0755}},
		{"directory user is empty", &Directory{"./test", "", "nobody", 0755}},
		{"directory group is empty", &Directory{"./test", "nobody", "", 0755}},
		{"directory permissions is empty", &Directory{"./test", "nobody", "nobody", 0}},
	}
	for _, tt := range tests {
		t.Run(tt.errMsg, func(t *testing.T) {
			err := (&FsPosix{}).ApplyDir(tt.conf)
			assert.Equal(t, tt.errMsg, err.Error())
		})
	}
}

func TestFsPosix_ApplyDir_EmptyDir(t *testing.T) {
	p := "./empty"
	fs := &FsPosix{}
	err := fs.ApplyDir(&Directory{
		Path: p,
		User: "nobody",
		Group: "nobody",
		Permissions: 0755,
	})
	assert.NoError(t, err)
	if assert.DirExists(t, p) {
		_, usr, grp, err := getDirInfo(p)
		assert.NoError(t, err, "can't get test dir info: %v", err)
		assert.Equal(t, "nobody", usr.Username)
		assert.Equal(t, "nobody", grp.Username)
	}
}

//func TestFsPosix_ApplyDir(t *testing.T) {
//	type args struct {
//		d *Directory
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr error
//	}{
//		{
//			name: "empty",
//			args: args {
//				d: &Directory{
//					Path:        "",
//					User:       "",
//					Group:       "",
//					Permissions: 0,
//				},
//			},
//			wantErr: nil,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			d := &FsPosix{}
//			err =
//			if err := d.ApplyDir(tt.args.d); (err != nil) != tt.wantErr {
//				t.Errorf("ApplyDir() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestFsPosix_ApplyFile(t *testing.T) {
//	type args struct {
//		f *File
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			d := &FsPosix{}
//			if err := d.ApplyFile(tt.args.f); (err != nil) != tt.wantErr {
//				t.Errorf("ApplyFile() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func getDirInfo(p string) (*syscall.Stat_t, *user.User, *user.User, error) {
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
