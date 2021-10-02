package yum

import (
	"context"
	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/shell/posix"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

//var execMock func(path string, args ...string) (*shell.Result, error)
//
//type shellMock struct{}
//
//func (u shellMock) Exec(_ context.Context, path string, args ...string) (*shell.Result, error) {
//	return execMock(path, args...)
//}

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
}

func TestYum_Init(t *testing.T) {
	sh := posix.New()
	mng := New(&packages.DriverConfig{}, sh)
	err := mng.Init(context.Background())
	assert.NoError(t, err)
}

func TestYum_RemovePackage(t *testing.T) {
	sh := posix.New()
	mng := New(&packages.DriverConfig{}, sh)
	assert.NoError(t, mng.RemovePackage(context.Background(), "epel-release"))
}

func TestYum_InstallPackage(t *testing.T) {
	sh := posix.New()
	mng := New(&packages.DriverConfig{}, sh)
	assert.NoError(t, mng.InstallPackage(context.Background(), "epel-release"))
}

func TestYum_InstallPackage_Repeat(t *testing.T) {
	sh := posix.New()
	mng := New(&packages.DriverConfig{}, sh)
	assert.NoError(t, mng.InstallPackage(context.Background(), "epel-release"))
	assert.NoError(t, mng.InstallPackage(context.Background(), "epel-release"))
}

//func TestYum_AddSource(t *testing.T) {
//	type fields struct {
//		conf *packages.DriverConfig
//	}
//	type args struct {
//		ctx context.Context
//		s   *packages.Source
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			y := Yum{
//				conf: tt.fields.conf,
//			}
//			if err := y.AddSource(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
//				t.Errorf("AddSource() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestYum_HostSetup(t *testing.T) {
//	type fields struct {
//		conf *packages.DriverConfig
//	}
//	type args struct {
//		ctx context.Context
//		c   *packages.DriverConfig
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			y := Yum{
//				conf: tt.fields.conf,
//			}
//			if err := y.HostSetup(tt.args.ctx, tt.args.c); (err != nil) != tt.wantErr {
//				t.Errorf("HostSetup() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestYum_InstallPackage(t *testing.T) {
//	type fields struct {
//		conf *packages.DriverConfig
//	}
//	type args struct {
//		ctx context.Context
//		p   *packages.Package
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			y := Yum{
//				conf: tt.fields.conf,
//			}
//			if err := y.InstallPackage(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("InstallPackage() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestYum_RemovePackage(t *testing.T) {
//	type fields struct {
//		conf *packages.DriverConfig
//	}
//	type args struct {
//		ctx context.Context
//		p   *packages.Package
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			y := Yum{
//				conf: tt.fields.conf,
//			}
//			if err := y.RemovePackage(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("RemovePackage() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestYum_RemoveSource(t *testing.T) {
//	type fields struct {
//		conf *packages.DriverConfig
//	}
//	type args struct {
//		ctx context.Context
//		s   *packages.Source
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			y := Yum{
//				conf: tt.fields.conf,
//			}
//			if err := y.RemoveSource(tt.args.ctx, tt.args.s); (err != nil) != tt.wantErr {
//				t.Errorf("RemoveSource() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
