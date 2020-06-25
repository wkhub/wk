package backends

import (
	"fmt"
	"strings"
)

var (
	backends        []Backend
	DefaultPrefixes = map[string]string{
		"gh":     "https://github.com/",
		"github": "https://github.com/",
		"gl":     "https://gitlab.com/",
		"gitlab": "https://gitlab.com/",
	}
)

type Backend interface {
	Name() string
	Match(source string) bool
	Fetch(source string) (string, error)
}

func Resolve(source string) (Backend, error) {
	source = UnPrefix(source)
	for _, backend := range backends {
		if backend.Match(source) {
			return backend, nil
		}
	}
	return nil, BackendError("Unknown source protocol")
}

func UnPrefix(source string) string {
	for prefix, replacement := range DefaultPrefixes {
		p := fmt.Sprintf(`%s:`, prefix)
		if strings.HasPrefix(source, p) {
			source = strings.Replace(source, p, replacement, 1)
			if !strings.HasSuffix(source, ".git") {
				source = source + ".git"
			}
			break
		}
	}
	return source
}

func Register(backend Backend) {
	backends = append(backends, backend)
}
