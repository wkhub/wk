package hooks

import (
	"github.com/wkhub/wk/shell"
)

type Hook interface {
	Match(session *shell.Session) bool
	Update(session *shell.Session) *shell.Session
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
func Execute(session *shell.Session) *shell.Session {
	// env := NewHookEnv()
	for _, hook := range hooks.hooks {
		if (*hook).Match(session) {
			(*hook).Update(session)
		}
	}
	return session
}

func register(hook Hook) {
	hooks.register(hook)
}
