package hooks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/shell"
)

const RBENVFILE = ".ruby-version"

func rbenvfile(path string) string {
	return filepath.Join(path, RBENVFILE)
}

type RubyHook BaseHook

func (h RubyHook) Match(session *shell.Session) bool {
	return fs.Exists(rbenvfile(session.Cwd))
}

func (h RubyHook) Update(session *shell.Session) *shell.Session {
	switch {
	case fs.Exists(rbenvfile(session.Cwd)):
		filename := rbenvfile(session.Cwd)
		version, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("Unable to read file %s", filename))
		}
		fmt.Println("version", version)
		// 	venv := venvFor(strings.Trim(string(name), " \n"))
		// 	cmd := fmt.Sprintf(". %s/bin/activate", venv)
		// 	env.Init = append(env.Init, cmd)
		// case fs.Exists(pipfile(path)):
		// 	cmd := fmt.Sprintf(". $(pipenv --venv)/bin/activate")
		// 	env.Init = append(env.Init, cmd)
		// case fs.Exists(venvdir(path)):
		// 	cmd := fmt.Sprintf(". %s/bin/activate", venvdir(path))
		// 	env.Init = append(env.Init, cmd)
	}
	return session
}

func init() {
	hooks.register(RubyHook{"ruby"})
}
