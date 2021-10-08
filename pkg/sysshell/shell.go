package sysshell

import (
	"context"
)

type Shell interface {
	Exec(ctx context.Context, path string, args ...string) (*Result, error)
}

type Result struct {
	Exit   int
	Stdout []string
	Stderr []string
}
