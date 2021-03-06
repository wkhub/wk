package shell

import (
	"os"
)

const SOURCE_BASHRC string = `
_mk_alias() {
	name=$1
	shift
	alias $name="_wk_eval $@"
	complete -o default -F _wk_$1 $name
}

_wk_eval() {
	. <(wk --bash --eval "$@")
}

_mk_alias wkon on
_mk_alias wknew new
_mk_alias wkcd cd
`

const BASH_EVAL string = `cd {{.Cwd}}
{{range .Env.Environ }}export {{ . }}
{{end}}

{{range .Init}}{{ . }}
{{end}}
`

type Bash struct {
	Name string
	Cmd  string
}

func (bash Bash) Run(session Session) {
	Command(bash.Cmd, session).Run()
}

func (bash Bash) Exec(session Session) int {
	session.AddCommand("exit $?")
	return RunWithExitCode(Command(bash.Cmd, session))
}

func (bash Bash) Rc() string {
	return SOURCE_BASHRC
}

func (bash Bash) Eval(session Session) {
	session.Render(BASH_EVAL, os.Stdout)
}

var BASH = Bash{"bash", "bash"}
