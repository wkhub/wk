package backends

import (
	"strings"
)

type HttpBackend struct {
}

func (b HttpBackend) Name() string {
	return "HTTP"
}

func (b HttpBackend) Match(source string) bool {
	return ((strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://")) &&
		!strings.HasSuffix(source, ".git"))
}

func (b HttpBackend) Fetch(source string) (string, error) {
	return "", nil
}

func init() {
	Register(HttpBackend{})
}
