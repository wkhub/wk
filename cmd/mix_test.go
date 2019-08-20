package cmd

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wkhub/wk/test"
)

func TestMixGuessArgs(t *testing.T) {
	cases := []struct {
		args   []string // Input args
		source string   // Expected source
		target string   // Expected target path
	}{
		{[]string{"test"}, "test", "{{.Cwd}}"},
		{[]string{"nested/test"}, "nested/test", "{{.Cwd}}"},
		{[]string{"/absolute/test"}, "/absolute/test", "{{.Cwd}}"},
		{[]string{"test", "."}, "test", "{{.Cwd}}"},
		{[]string{"test", "target/path"}, "test", "{{.Cwd}}/target/path"},
		{[]string{"test", "/target/path"}, "test", "/target/path"},
	}

	for _, cc := range cases {
		test.Sandbox(func(ctx test.SCtx) {
			testName := strings.Join(cc.args, "+")
			t.Run(testName, func(t *testing.T) {
				source, target := parseMixArgs(cc.args)
				assert.Equal(t, ctx.Render(filepath.FromSlash(cc.source)), source,
					"Wrong parsed source")
				assert.Equal(t, ctx.Render(filepath.FromSlash(cc.target)), target,
					"Wrong parsed target path")
			})
		})
	}
}
