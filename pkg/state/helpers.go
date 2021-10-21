package state

import (
	"io/fs"
)

func BoolRef(b bool) *bool {
	return &b
}

func BoolDeref(b *bool) bool {
	if nil == b {
		return false
	}
	return *b
}

func StringRef(s string) *string {
	return &s
}

func StringDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func FileModeDeref(i *fs.FileMode) fs.FileMode {
	if i == nil {
		return 0
	}
	return *i
}
