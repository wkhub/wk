package hooks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/wkhub/wk/fs"
)

const RBENVFILE = ".ruby-version"

func rbenvfile(path string) string {
	return filepath.Join(path, RBENVFILE)
}

type RubyHook BaseHook

func (h RubyHook) Match(path string) bool {
	return fs.Exists(rbenvfile(path))
}

func (h RubyHook) GetEnv(path string) HookEnv {
	env := NewHookEnv()
	switch {
	case fs.Exists(rbenvfile(path)):
		version, err := ioutil.ReadFile(rbenvfile(path))
		if err != nil {
			panic(fmt.Sprintf("Unable to read file %s", path))
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
	return env
}

func init() {
	hooks.register(RubyHook{"ruby"})
}
