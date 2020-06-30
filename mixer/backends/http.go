package backends

import (
	"strings"

	"github.com/pkg/errors"
)

type HTTPBackend struct {
}

func (b HTTPBackend) Name() string {
	return "HTTP"
}

func (b HTTPBackend) Match(source string) bool {
	return ((strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://")) &&
		!strings.HasSuffix(source, ".git"))
}

func (b HTTPBackend) Fetch(source string) (string, error) {
	return "", errors.New("Not implemented")
}

func init() {
	Register(HTTPBackend{})
}
