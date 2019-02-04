package backends

import "strings"

type GistBackend struct {
}

func (b GistBackend) Name() string {
	return "Git"
}

func (b GistBackend) Match(source string) bool {
	return strings.HasPrefix(source, "gist:")
}

func (b GistBackend) Fetch(source string) (string, error) {
	return "", nil
}

func init() {
	Register(GistBackend{})
}
