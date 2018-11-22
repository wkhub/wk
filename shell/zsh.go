package shell

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/wkhub/wk/fs"
)

const ZSHRC string = `source ~/.zshrc
eval $WK_ZSH_INIT
export ZDOTDIR=$WK_ZDOTDIR_ORIG
`

const SOURCE_ZSHRC string = `
_wk_eval() {
    # echo "Command is wk --zsh --eval $@"
	. <(wk --zsh --eval "$@")
}

_wk_projects() {
	reply=( $(wk list) )
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

const ZSH_EVAL string = `cd {{.Cwd}}
{{range .Env.Environ }}export {{ . }}
{{end}}

{{range .Init}}{{ . }}
{{end}}
`

type Zsh struct {
	ShellHelper
	Cmd string
}

func (zsh Zsh) buildZDotenv() {
	zsh.ensureConfigDir()
	zshrc := zsh.configFile(".zshrc")
	if !fs.Exists(zshrc) {
		ioutil.WriteFile(zshrc, []byte(ZSHRC), 0655)
	}
}

func (zsh Zsh) Run(session Session) {
	zsh.buildZDotenv()
	session.Env["ZDOTDIR"] = zsh.configDir()
	session.Env["WK_ZSH_INIT"] = strings.Join(session.Init, "\n")

	shell := exec.Command(zsh.Cmd)
	shell.Env = session.Env.Environ()
	shell.Dir = session.Cwd
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Run()
}

func (zsh Zsh) Eval(session Session) {
	session.Render(ZSH_EVAL, os.Stdout)
}

func (zsh Zsh) Rc() string {
	return SOURCE_ZSHRC
}

var ZSH = Zsh{ShellHelper{"zsh"}, "zsh"}
