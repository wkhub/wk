package hooks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/shell"
)

const PIPFILE = "Pipfile"
const VENVFILE = ".venv"
const VENVDIR = "venv"

func pipfile(path string) string {
	return filepath.Join(path, PIPFILE)
}

func venvfile(path string) string {
	return filepath.Join(path, VENVFILE)
}

func venvdir(path string) string {
	return filepath.Join(path, VENVDIR)
}

func WorkonHome() string {
	return shell.GetEnv().Get("WORKON_HOME", "~/.virtualenvs")
}

func venvFor(name string) string {
	return filepath.Join(WorkonHome(), name)
}

type PythonHook BaseHook

func (h PythonHook) Match(path string) bool {
	return fs.Exists(venvfile(path)) || fs.Exists(pipfile(path)) || fs.Exists(venvdir(path))
}

func (h PythonHook) GetEnv(path string) HookEnv {
	env := NewHookEnv()
	switch {
	case fs.Exists(venvfile(path)):
		name, err := ioutil.ReadFile(venvfile(path))
		if err != nil {
			panic(fmt.Sprintf("Unable to read file %s", path))
		}
		venv := venvFor(strings.Trim(string(name), " \n"))
		cmd := fmt.Sprintf(". %s/bin/activate", venv)
		env.Init = append(env.Init, cmd)
	case fs.Exists(pipfile(path)):
		cmd := fmt.Sprintf(". $(pipenv --venv)/bin/activate")
		env.Init = append(env.Init, cmd)
	case fs.Exists(venvdir(path)):
		cmd := fmt.Sprintf(". %s/bin/activate", venvdir(path))
		env.Init = append(env.Init, cmd)
	}
	return env
}

func init() {
	hooks.register(PythonHook{"python"})
}
