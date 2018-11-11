package hooks

import (
	"github.com/wkhub/wk/shell"
)

type HookEnv struct {
	Env  shell.Env // Exported environment variables
	Init []string  // Commands to run on entering the workspace
	// Dirs		map[string]string	// Directories shortcuts
	// Commands	map[string]string	// Commands shortcuts
}

func (he *HookEnv) Merge(other HookEnv) {
	he.Env.Update(other.Env)
	he.Init = append(he.Init, other.Init...)
}

func NewHookEnv() HookEnv {
	return HookEnv{
		Env:  make(shell.Env),
		Init: []string{},
	}
}

type Hook interface {
	Match(path string) bool
	GetEnv(path string) HookEnv
}

type BaseHook struct {
	Name string
}

type Hooks struct {
	hooks []*Hook
}

func (h *Hooks) register(hook Hook) {
	h.hooks = append(h.hooks, &hook)
}

var hooks = Hooks{[]*Hook{}}

// Execute will find hooks for a given directory
// and execute them
func Execute(path string) HookEnv {
	env := NewHookEnv()
	for _, hook := range hooks.hooks {
		if (*hook).Match(path) {
			env.Merge((*hook).GetEnv(path))
		}
	}
	return env
}

func register(hook Hook) {
	hooks.register(hook)
}
