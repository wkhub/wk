package hooks

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/shell"
)

const PIPFILE = "Pipfile"
const VENV = "venv"
const DOTVENV = ".venv"

func venvFor(name string) string {
	return filepath.Join(shell.GetEnv().Get("WORKON_HOME", "~/.virtualenvs"), name)
}

type PythonHook BaseHook

// Match if one of the known file is found
func (h PythonHook) Match(s *shell.Session) bool {
	for _, filename := range []string{VENV, DOTVENV, PIPFILE} {
		if fs.Exists(filepath.Join(s.Cwd, filename)) {
			return true
		}
	}
	return false
}

// Update activate virtualenv and export some aliases
func (h PythonHook) Update(session *shell.Session) *shell.Session {
	var venv string
	candidates := []string{VENV, DOTVENV}
	for _, filename := range candidates {
		path := filepath.Join(session.Cwd, filename)
		if fs.Exists(path) {
			venv = extractVenv(path)
			break
		}
	}
	if venv == "" && fs.Exists(filepath.Join(session.Cwd, PIPFILE)) {
		cmd := exec.Command("pipenv", "--venv")
		out, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		venv = string(out)
	}
	if venv != "" {
		session.Init = append(session.Init, fmt.Sprintf("echo 'Using virtualenv %s'", venv))
		session.Init = append(session.Init, fmt.Sprintf(". %s/bin/activate", venv))
		session.Dirs["virtualenv"] = venv
		session.Dirs["venv"] = venv
	}
	return session
}

func extractVenv(path string) string {
	switch {
	case fs.IsFile(path):
		name, err := ioutil.ReadFile(path)
		if err != nil {
			panic(fmt.Sprintf("Unable to read file %s", path))
		}
		return venvFor(strings.Trim(string(name), " \n"))
	default: // It's a directory
		return path
	}
}

func init() {
	hooks.register(PythonHook{"python"})
}
