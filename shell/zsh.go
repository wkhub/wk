package shell

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/wkhub/wk/fs"
)

const tmpZshrc string = `source ~/.zshrc
eval $WK_ZSH_INIT
export ZDOTDIR=$WK_ZDOTDIR_ORIG
`

const srcZshrc string = `
_wk_eval() {
    # echo "Command is wk --zsh --eval $@"
	. <(wk --zsh --eval "$@")
}

_wk_projects() {
	# reply=( $(wk list) );
	reply=( ${(f)$(wk list)} );
}

_wk_cd_dirs() {
	reply=( $(wk cd --list) )
}

_mk_alias() {
	name=$1
	shift
	alias $name="_wk_eval $@"
	compdef "_wk $@" $name	
}

_mk_alias wknew new
_mk_alias wkon on
_mk_alias wkcd cd

compctl -K _wk_projects wkon
`

const zshEval string = `cd {{.Cwd}}
{{range .Env.Environ }}export {{ . }}
{{end}}

{{range .Init}}{{ . }}
{{end}}
`

type Zsh struct {
	Helper
	Cmd string
}

func (zsh Zsh) buildZDotenv() {
	zsh.ensureConfigDir()
	zshrc := zsh.configFile(".zshrc")
	if !fs.Exists(zshrc) {
		if err := ioutil.WriteFile(zshrc, []byte(tmpZshrc), 0655); err != nil {
			panic(err)
		}
	}
}

func (zsh Zsh) buildCommand(session Session) *exec.Cmd {
	zsh.buildZDotenv()
	session.Env["ZDOTDIR"] = zsh.configDir()
	session.Env["WK_ZSH_INIT"] = strings.Join(session.Init, "\n")

	return Command(zsh.Cmd, session)
}

func (zsh Zsh) Run(session Session) {
	if err := zsh.buildCommand(session).Run(); err != nil {
		panic(err)
	}
}

func (zsh Zsh) Exec(session Session) int {
	session.AddCommand("exit $?")
	return RunWithExitCode(zsh.buildCommand(session))
}

func (zsh Zsh) Eval(session Session) {
	session.Render(zshEval, os.Stdout)
}

func (zsh Zsh) Rc() string {
	return srcZshrc
}

var ZSH = Zsh{Helper{"zsh"}, "zsh"}
