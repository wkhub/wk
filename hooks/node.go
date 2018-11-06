package hooks

import (
	"fmt"
	"os"
	"path/filepath"
)

const NVMRC = ".nvmrc"

type NvmHook BaseHook

func (h NvmHook) Match(path string) bool {
	path = filepath.Join(path, NVMRC)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (h NvmHook) GetEnv(path string) ([]string, []string) {
	cmd := fmt.Sprintf("nvm use $(cat %s/.nvmrc)", path)
	return []string{}, []string{cmd}
}

func init() {
	hooks.register(NvmHook{"nvm"})
}
