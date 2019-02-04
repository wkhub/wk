package backends

import (
	"github.com/wkhub/wk/fs"
)

type FileBackend struct {
}

func (b FileBackend) Name() string {
	return "File"
}

func (b FileBackend) Match(source string) bool {
	return fs.Exists(source) && fs.IsDir(source)
}

func (b FileBackend) Fetch(source string) (string, error) {
	return source, nil
}

func init() {
	Register(FileBackend{})
}
