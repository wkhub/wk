package mixer

import (
	"bytes"
	"html/template"
)

type Context map[string]interface{}

func (ctx Context) Render(txt string) string {
	var out bytes.Buffer
	tmpl, err := template.New("test").Parse(txt)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&out, ctx)
	if err != nil {
		panic(err)
	}
	return out.String()
}
