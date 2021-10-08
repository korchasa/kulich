package posix

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/korchasa/ruchki/pkg/sysshell"
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

func (p *Posix) Exec(ctx context.Context, path string, args ...string) (*sysshell.Result, error) {
	log.Debugf("Shell exec `%s %s`", path, strings.Join(args, " "))
	res := &sysshell.Result{}
	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			res.Exit = e.ExitCode()
		} else {
			return nil, fmt.Errorf("can't exec `%s %s`: %w", path, strings.Join(args, " "), err)
		}
	}

	outStr, errStr := stdout.String(), stdout.String()
	res.Stdout = strings.Split(outStr, "\n")
	res.Stderr = strings.Split(errStr, "\n")

	p.History = append(p.History, HistoryExec{
		Path:   path,
		Args:   args,
		Result: res,
	})

	log.Debugf(`Shell exec complete with code "%d"\n
### stdout ##################\n
%s\n
### stderr ##################\n
%s\n`, res.Exit, outStr, errStr)
	return res, nil
}

type HistoryExec struct {
	Path   string
	Args   []string
	Result *sysshell.Result
}
