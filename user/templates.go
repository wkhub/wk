package user

import "path/filepath"

const _TEMPLATES_DIR = "templates"

func (h Home) TemplatesDir() string {
	return filepath.Join(h.Path, _TEMPLATES_DIR)
}
