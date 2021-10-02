package posix

import (
	"bytes"
	"context"
	"fmt"
	"github.com/korchasa/ruchki/pkg/shell"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Posix struct {
	History []HistoryExec
}

func New() *Posix {
	return &Posix{}
}

func (p *Posix) Exec(ctx context.Context, path string, args ...string) (*shell.Result, error) {
	log.Infof("Shell exec `%s %s`", path, strings.Join(args, " "))
	res := &shell.Result{}
	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			res.Exit = exitError.ExitCode()
		} else {
			return nil, fmt.Errorf("can't exec `%s %s`: %v", path, strings.Join(args, " "), err)
		}
	}

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	res.Stdout = strings.Split(outStr, "\n")
	res.Stderr = strings.Split(errStr, "\n")

	p.History = append(p.History, HistoryExec{
		Path:   path,
		Args:   args,
		Result: res,
	})

	log.Debugf("Shell exec complete with code `%d`\nstdout: %s\n stderr: %s", res.Exit, outStr, errStr)
	return res, nil
}

type HistoryExec struct {
	Path   string
	Args   []string
	Result *shell.Result
}
