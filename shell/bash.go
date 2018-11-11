package shell

import (
	"os"
	"os/exec"
)

type Bash struct {
	Name string
	Cmd  string
}

func (bash Bash) Run(cwd string, env []string, cmds []string) {
	shell := exec.Command(bash.Cmd)
	shell.Dir = cwd
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Run()
}

func (bash Bash) Rc() string {
	return ""
}

func (bash Bash) Eval(cwd string, env []string, cmds []string) string {
	return ""
}

var BASH = Bash{"bash", "bash"}
