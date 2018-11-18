package hooks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const ENVFILE = ".env"

type EnvHook BaseHook

func (h EnvHook) Match(path string) bool {
	path = filepath.Join(path, ENVFILE)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (h EnvHook) GetEnv(path string) HookEnv {
	env := NewHookEnv()
	filename := filepath.Join(path, ENVFILE)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file %s", filename))
	}
	for _, line := range strings.Split(string(text), "\n") {
		if strings.TrimSpace(line) != "" {
			parts := strings.Split(line, "=")
			env.Env[parts[0]] = parts[1]
		}
	}
	return env
}

func init() {
	hooks.register(EnvHook{"env"})
}
