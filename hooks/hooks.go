package hooks

type Hook interface {
	Match(path string) bool
	GetEnv(path string) ([]string, []string)
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
func Execute(path string) ([]string, []string) {
	envs, scripts := []string{}, []string{}
	for _, hook := range hooks.hooks {
		if (*hook).Match(path) {
			env, script := (*hook).GetEnv(path)
			envs = append(envs, env...)
			scripts = append(scripts, script...)
		}
	}
	return envs, scripts
}

func register(hook Hook) {
	hooks.register(hook)
}
