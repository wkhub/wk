package mixer

import (
	"bytes"
	"html/template"
	"strings"
)

type Context map[string]interface{}

var helpers = template.FuncMap{
	"replace": func(value interface{}, replacement interface{}, str interface{}) string {
		return strings.Replace(str.(string), value.(string), replacement.(string), -1)
	},
	"lower": strings.ToLower,
	"title": strings.ToTitle,
	"contains": func(key string, slice []string) bool {
		for _, value := range slice {
			if value == key {
				return true
			}
		}
		return false
	},
}

// Render a template string using the current context
func (ctx Context) Render(txt string) string {
	var out bytes.Buffer
	tmpl, err := template.New("test").Funcs(helpers).Parse(txt)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&out, map[string]interface{}{
		"ctx": ctx,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}

// RenderList render each string using the current context
func (ctx Context) RenderList(list []string) []string {
	out := []string{}
	for _, pattern := range list {
		out = append(out, ctx.Render(pattern))
	}
	return out
}
