package cmd

import (
	"bytes"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/wkhub/wk/test"
)

func render(txt string, ctx test.SCtx) string {
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

func TestGuessArgs(t *testing.T) {
	cases := []struct {
		args []string // Input args
		name string   // Expected name
		path string   // Expected path
	}{
		{[]string{"test"}, "test", "test"},          // Project name as only arg
		{[]string{"test", "."}, "test", "{{.Cwd}}"}, // Project name as only arg
		{[]string{"."}, "{{.Dirname}}", "{{.Cwd}}"}, // Project name as only arg
	}

	for _, cc := range cases {
		test.Sandbox(func(ctx test.SCtx) {
			testName := strings.Join(cc.args, "+")
			t.Run(testName, func(t *testing.T) {
				name, path := newGuessArgs(cc.args)
				assert.Equal(t, render(cc.name, ctx), name, "Wrong name")
				assert.Equal(t, render(cc.path, ctx), path, "Wrong path")
			})
		})
	}
}
