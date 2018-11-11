package shell

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/wkhub/wk/fs"
)

const ZSHRC string = `source ~/.zshrc
eval $WK_ZSH_INIT
export ZDOTDIR=$WK_ZDOTDIR_ORIG
`

const SOURCE_ZSHRC string = `
_mk_alias() {
	name=$1
	shift
	alias $name="_wk_and_source $@"
	compdef '_wk $@' $name	
}

_wk_and_source() {
	#echo "Command is wk --zsh --eval $@"
	. <(wk --zsh --eval "$@")
}

_mk_alias goon on
_mk_alias gonew new
`

const ZSH_EVAL string = `cd {{.Cwd}}
{{range .Env}}export {{ . }}
{{end}}

{{range .Commands}}{{ . }}
{{end}}
`

type EvalCtx struct {
	Cwd      string
	Env      []string
	Commands []string
}

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

func (zsh Zsh) Run(cwd string, env []string, cmds []string) {
	zsh.buildZDotenv()
	newEnv := append(os.Environ(), env...)

	newEnv = append(newEnv, "ZDOTDIR="+zsh.configDir())
	newEnv = append(newEnv, "WK_ZSH_INIT="+strings.Join(cmds, "\n"))

	shell := exec.Command(zsh.Cmd)
	shell.Env = newEnv
	shell.Dir = cwd
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Run()
}

func (zsh Zsh) Eval(cwd string, env []string, cmds []string) string {
	var out bytes.Buffer
	tmpl, err := template.New("zsh-eval").Parse(ZSH_EVAL)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&out, EvalCtx{cwd, env, cmds})
	if err != nil {
		panic(err)
	}
	return out.String()
}

func (zsh Zsh) Rc() string {
	return SOURCE_ZSHRC
}

var ZSH = Zsh{ShellHelper{"zsh"}, "zsh"}
