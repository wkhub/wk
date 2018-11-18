package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wkhub/wk/test"
)

func TestNewGuessArgs(t *testing.T) {
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
				assert.Equal(t, ctx.Render(cc.name), name, "Wrong name")
				assert.Equal(t, ctx.Render(cc.path), path, "Wrong path")
			})
		})
	}
}
