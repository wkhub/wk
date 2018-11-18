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

func (h PythonHook) Match(s *shell.Session) bool {
	return fs.Exists(venvfile(s.Cwd)) || fs.Exists(pipfile(s.Cwd)) || fs.Exists(venvdir(s.Cwd))
}

func (h PythonHook) Update(session *shell.Session) *shell.Session {
	switch {
	case fs.Exists(venvfile(session.Cwd)):
		filename := venvfile(session.Cwd)
		name, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("Unable to read file %s", filename))
		}
		venv := venvFor(strings.Trim(string(name), " \n"))
		cmd := fmt.Sprintf(". %s/bin/activate", venv)
		session.Init = append(session.Init, cmd)
	case fs.Exists(pipfile(session.Cwd)):
		cmd := fmt.Sprintf(". $(pipenv --venv)/bin/activate")
		session.Init = append(session.Init, cmd)
	case fs.Exists(venvdir(session.Cwd)):
		cmd := fmt.Sprintf(". %s/bin/activate", venvdir(session.Cwd))
		session.Init = append(session.Init, cmd)
	}
	return session
}

func init() {
	hooks.register(PythonHook{"python"})
}
