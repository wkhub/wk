package shell

import (
	"os"
	"path/filepath"

	"github.com/wkhub/wk/fs"
)

func Current() Shell {
	// out, _ := exec.Command("ps", "-p $$", "-ocomm=").Output()
	// shell := string(out[:])
	// fmt.Print(os.Getppid())
	// shell := os.Getenv("SHELL")
	return ZSH
}

type Shell interface {
	Run(cwd string, env []string, cmds []string)
}

type ShellHelper struct {
	Name string
}

func (sh ShellHelper) configDir() string {
	wkHome := fs.Home()
	return filepath.Join(wkHome, "shells", sh.Name)
}

func (sh ShellHelper) ensureConfigDir() {
	os.MkdirAll(sh.configDir(), 0755)
}

func (sh ShellHelper) configFile(name string) string {
	return filepath.Join(sh.configDir(), name)
}
