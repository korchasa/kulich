package posix

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/korchasa/kulich/pkg/sysshell"
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

func (p *Posix) Exec(cmd *exec.Cmd) (*sysshell.Result, error) {
	log.Debugf("Shell exec `%s %s`", cmd.Path, strings.Join(cmd.Args, " "))
	res := &sysshell.Result{}
	var stdout, stderr bytes.Buffer

	//cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		var e *exec.ExitError
		if errors.As(err, &e) {
			res.Exit = e.ExitCode()
		} else {
			return nil, fmt.Errorf("can't exec `%s %s`: %w", cmd.Path, strings.Join(cmd.Args, " "), err)
		}
	}

	outStr, errStr := stdout.String(), stderr.String()
	res.Stdout = strings.Split(outStr, "\n")
	res.Stderr = strings.Split(errStr, "\n")

	p.History = append(p.History, HistoryExec{
		Command: cmd,
		Result:  res,
	})

	log.Debugf(
		`Shell exec complete with code "%d"
### stdout ##################
%s
### stderr ##################
%s
#############################`,
		res.Exit,
		stdout.String(),
		stderr.String())
	return res, nil
}

func (p *Posix) SafeExec(command string) ([]string, error) {
	parts := strings.Split(command, " ")
	cmd := exec.Command(parts[0], parts[1:]...)
	res, err := p.Exec(cmd)
	if err != nil {
		return []string{}, err
	}
	if res.Exit != 0 {
		return res.Stdout,
			fmt.Errorf(
				"non-zero exit code (%d) from `%s`:\nstdout: %s\nstderr: %s",
				res.Exit,
				cmd.Path,
				strings.Join(res.Stdout, "\n"),
				strings.Join(res.Stderr, "\n"))
	}
	return res.Stdout, nil
}

func (p *Posix) SafeExecf(command string, args ...interface{}) ([]string, error) {
	return p.SafeExec(fmt.Sprintf(command, args...))
}

type HistoryExec struct {
	Command *exec.Cmd
	Result  *sysshell.Result
}
