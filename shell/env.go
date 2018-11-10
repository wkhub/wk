package shell

import (
	"os"
	"strconv"
	"strings"
)

// Env is a map equivalent of os.Environ
type Env map[string]string

// GetEnv returns a Env populated with the actual environment variables
func GetEnv() Env {
	env := make(Env)

	for _, kv := range os.Environ() {
		splitted := strings.SplitN(kv, "=", 2)
		env[splitted[0]] = splitted[1]
	}

	return env
}

// Get try to fetch an envionment variable and if not found return a fallback value
func (e Env) Get(key string, fallback string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	} else {
		return fallback
	}
}

// Has checks wether or not an environment exists
func (e Env) Has(key string) bool {
	_, found := os.LookupEnv(key)
	return found
}

// GetInt try to fetch an envionment variable and if not found return a fallback value
// The value is casted as an int.
// If the cast is not possible, it will fail silently and return the fallback value
func (e Env) GetInt(key string, fallback int) int {
	if e.Has(key) {
		value, err := strconv.Atoi(e[key])
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}

// Clone returns a new Env with same values
func (e Env) Clone() Env {
	// Create the target map
	new := make(Env)

	// Copy from the original map to the target map
	for key, value := range e {
		new[key] = value
	}
	return new
}

// Environ convert the Env into a []string compatible with os.Environ
func (e Env) Environ() []string {
	environ := []string{}
	for key, value := range e {
		environ = append(environ, strings.Join([]string{key, value}, "="))
	}
	return environ
}
