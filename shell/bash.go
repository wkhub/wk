package shell

import (
	"os"
	"os/exec"
)

const SOURCE_BASHRC string = `
_mk_alias() {
	name=$1
	shift
	alias $name="_wk_and_source $@"
	complete -o default -F _wk_$1 $name
}

_wk_and_source() {
	#echo "Command is wk --bash --eval $@"
	. <(wk --bash --eval "$@")
}

_mk_alias wkon on
_mk_alias wknew new
_mk_alias wkcd cd
`

const BASH_EVAL string = `cd {{.Cwd}}
{{range .Env}}export {{ . }}
{{end}}

{{range .Commands}}{{ . }}
{{end}}
`

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
	return SOURCE_BASHRC
}

func (bash Bash) Eval(cwd string, env []string, cmds []string) string {
	return ""
}

var BASH = Bash{"bash", "bash"}
