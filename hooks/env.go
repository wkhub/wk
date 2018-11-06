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

func (h EnvHook) GetEnv(path string) ([]string, []string) {
	filename := filepath.Join(path, ENVFILE)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file %s", filename))
	}
	env := strings.Split(string(text), "\n")
	return env, []string{}
}

func init() {
	hooks.register(EnvHook{"env"})
}
