package backends

import (
	"strings"
)

var backends []Backend

type Backend interface {
	Name() string
	Match(source string) bool
	Fetch(source string) (string, error)
}

func Resolve(source string) (Backend, error) {
	source = Unalias(source)
	for _, backend := range backends {
		if backend.Match(source) {
			return backend, nil
		}
	}
	return nil, BackendError("Unknown source protocol")
}

func Unalias(source string) string {
	if strings.HasPrefix(source, "gh:") {
		source = strings.Replace(source, "gh:", "https://github.com/", 1)
		if !strings.HasSuffix(source, ".git") {
			source = source + ".git"
		}
	}
	return source
}

func Register(backend Backend) {
	backends = append(backends, backend)
}
