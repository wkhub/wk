package hooks

import (
	"fmt"
	"path/filepath"

	"github.com/wkhub/wk/fs"
)

const PIPFILE = "Pipfile"
const VENVFILE = ".venv"
const VENVDIR = "venv"

type PythonHook BaseHook

func (h PythonHook) Match(path string) bool {
	return fs.Exists(filepath.Join(path, VENVFILE)) || fs.Exists(filepath.Join(path, PIPFILE)) || fs.Exists(filepath.Join(path, VENVFILE))
}

func (h PythonHook) GetEnv(path string) ([]string, []string) {
	env := []string{}
	cmds := []string{}
	switch {
	case fs.Exists(filepath.Join(path, VENVFILE)):
		cmd := fmt.Sprintf("workon -n $(cat %s/.venv)", path)
		cmds = append(cmds, cmd)
	case fs.Exists(filepath.Join(path, PIPFILE)):
		cmd := fmt.Sprintf(". $(pipenv --venv)/bin/activate")
		cmds = append(cmds, cmd)
	case fs.Exists(filepath.Join(path, VENVFILE)):
		cmd := fmt.Sprintf(". venv/bin/activate")
		cmds = append(cmds, cmd)
	}
	return env, cmds
}

func init() {
	hooks.register(PythonHook{"python"})
}
