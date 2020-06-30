package user

import (
	"path/filepath"

	"github.com/wkhub/wk/shell"
)

const _shellsDir = "shells"

func (h Home) ShellsDir() string {
	return filepath.Join(h.Path, _shellsDir)
}

func (user User) Shell() shell.Shell {
	return shell.ZSH
}
