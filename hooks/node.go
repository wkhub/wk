package hooks

import (
	"fmt"
	"path/filepath"

	"github.com/wkhub/wk/fs"

	"github.com/wkhub/wk/shell"
)

const NVMRC = ".nvmrc"

type NvmHook BaseHook

func (h NvmHook) Match(session *shell.Session) bool {
	path := filepath.Join(session.Cwd, NVMRC)
	return fs.Exists(path)
}

func (h NvmHook) Update(session *shell.Session) *shell.Session {
	cmd := fmt.Sprintf("nvm use $(cat %s/.nvmrc)", session.Cwd)
	session.Init = append(session.Init, cmd)
	return session
}

func init() {
	hooks.register(NvmHook{"nvm"})
}
