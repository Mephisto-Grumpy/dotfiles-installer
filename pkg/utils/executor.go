package utils

import (
	"os"
	"os/exec"
)

type Executor interface {
	RunCmd(string, ...string) error
}

type CmdExecutor struct{}

func (e *CmdExecutor) RunCmd(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
