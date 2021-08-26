package ruchki

import "testing"

func TestFsCommon_Setup(t *testing.T) {
	type args struct {
		c *FsDriverConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &FsCommon{}
			if err := d.Setup(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsCommon_ApplyDir(t *testing.T) {
	type args struct {
		d *Directory
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &FsCommon{}
			if err := d.ApplyDir(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("ApplyDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFsCommon_ApplyFile(t *testing.T) {
	type args struct {
		f *File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &FsCommon{}
			if err := d.ApplyFile(tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("ApplyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
