package user

import "path/filepath"

const _templatesDir = "templates"

func (h Home) TemplatesDir() string {
	return filepath.Join(h.Path, _templatesDir)
}
