package yum_test

import (
	"context"
	"github.com/korchasa/ruchki/pkg/packages/yum"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/korchasa/ruchki/pkg/packages"
	"github.com/korchasa/ruchki/pkg/sysshell/posix"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

//var execMock func(path string, args ...string) (*shell.Result, error)
//
//type shellMock struct{}
//
//func (u shellMock) Exec(_ context.Context, path string, args ...string) (*shell.Result, error) {
//	return execMock(path, args...)
//}

type YumIntegrationTestSuite struct {
	suite.Suite
}

func (suite *YumIntegrationTestSuite) SetupTest() {
	// log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
}

func TestYumIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(YumIntegrationTestSuite))
}

func (suite *YumIntegrationTestSuite) TestYum_Init() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	err := mng.Init(context.Background())
	assert.NoError(suite.T(), err)
}

func (suite *YumIntegrationTestSuite) TestYum_RemovePackage() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.RemovePackage(context.Background(), "epel-release"))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
}

func (suite *YumIntegrationTestSuite) TestYum_InstallPackage_Repeat() {
	sh := posix.New()
	mng := yum.New(&packages.DriverConfig{}, sh)
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
	assert.NoError(suite.T(), mng.InstallPackage(context.Background(), "epel-release"))
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
