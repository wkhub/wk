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

var ZSH = Zsh{ShellHelper{"zsh"}, "zsh"}
