package user

import (
	"path/filepath"

	"github.com/wkhub/wk/shell"
)

const _SHELLS_DIR = "shells"

func (h Home) ShellsDir() string {
	return filepath.Join(h.Path, _SHELLS_DIR)
}

func (user User) Shell() shell.Shell {
	return shell.ZSH
}
