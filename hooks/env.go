package hooks

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/shell"
)

const ENVFILE = ".env"

type EnvHook BaseHook

func (h EnvHook) Match(session *shell.Session) bool {
	path := filepath.Join(session.Cwd, ENVFILE)
	return fs.Exists(path)
}

func (h EnvHook) Update(session *shell.Session) *shell.Session {
	filename := filepath.Join(session.Cwd, ENVFILE)
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to read file %s", filename))
	}
	for _, line := range strings.Split(string(text), "\n") {
		if strings.TrimSpace(line) != "" {
			parts := strings.Split(line, "=")
			session.Env[parts[0]] = parts[1]
		}
	}
	return session
}

func init() {
	hooks.register(EnvHook{"env"})
}
