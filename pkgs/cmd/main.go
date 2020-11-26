package cmd

import (
	"os/exec"
)

// Cmd is an interface
type Cmd interface {
	Execute() (result, error)
}

type cmd struct {
	command string
	args    []string
}

type result struct {
	Data     string
	ExitCode int
}

func (c cmd) Execute() (result, error) {
	out, err := exec.Command(c.command, c.args...).Output()
	if err != nil {
		return result{
			Data:     "",
			ExitCode: 2,
		}, err
	}
	return result{
		Data:     string(out),
		ExitCode: 0,
	}, nil
}

// NewCmd returns a Cmd interface
func NewCmd(command string, args []string) Cmd {
	return cmd{
		command: command,
		args:    args,
	}
}
