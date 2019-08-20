package mixer

import (
	"github.com/noirbizarre/gonja"
)

type Context map[string]interface{}

// Render a template string using the current context
func (ctx Context) Render(txt string) (string, error) {
	tpl, err := gonja.FromString(txt)
	if err != nil {
		return "", err
	}
	// Now you can render the template with the given
	// gonja.Context how often you want to.
	out, err := tpl.Execute(ctx)
	if err != nil {
		return "", err
	}
	return out, nil
}

// RenderList render each string using the current context
func (ctx Context) RenderList(list []string) ([]string, error) {
	out := []string{}
	for _, pattern := range list {
		rendered, err := ctx.Render(pattern)
		if err != nil {
			return out, err
		}
		out = append(out, rendered)
	}
	return out, nil
}
