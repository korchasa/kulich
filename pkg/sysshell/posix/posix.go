package posix

import (
	"bytes"
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

	outStr, errStr := stdout.String(), stdout.String()
	res.Stdout = strings.Split(outStr, "\n")
	res.Stderr = strings.Split(errStr, "\n")

	p.History = append(p.History, HistoryExec{
		Command: cmd,
		Result:  res,
	})

	log.Debugf(`Shell exec complete with code "%d"\n
### stdout ##################\n
%s\n
### stderr ##################\n
%s\n`, res.Exit, outStr, errStr)
	return res, nil
}

func (p *Posix) SafeExec(command string) ([]string, error) {
	parts := strings.Split(command, " ")
	res, err := p.Exec(&exec.Cmd{
		Path:         parts[0],
		Args:         parts,
		Env:          nil,
		Dir:          "",
		Stdin:        nil,
		Stdout:       nil,
		Stderr:       nil,
		ExtraFiles:   nil,
		SysProcAttr:  nil,
		Process:      nil,
		ProcessState: nil,
	})
	if err != nil {
		return res.Stdout, err
	}
	if res.Exit != 0 {
		return res.Stdout, fmt.Errorf("non-zero exit code (%d) from %s: %s", res.Exit, parts[0], strings.Join(res.Stderr, "\n"))
	}
	return res.Stdout, nil
}

type HistoryExec struct {
	Command *exec.Cmd
	Result  *sysshell.Result
}
