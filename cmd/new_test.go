package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

type Ctx struct {
	Cwd     string // Current working directoy
	Dirname string // dirname of the current working directory
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Sandbox(cb func(ctx Ctx)) {
	cwd, err := os.Getwd()
	handleError(err)
	tmp, err := ioutil.TempDir("", "")
	handleError(err)
	os.Chdir(tmp)
	ctx := Ctx{tmp, filepath.Base(tmp)}

	defer func() {
		os.Chdir(cwd)
	}()

	cb(ctx)
}

func render(txt string, ctx Ctx) string {
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
		Sandbox(func(ctx Ctx) {
			name, path := newGuessArgs(cc.args)
			assert.Equal(t, render(cc.name, ctx), name, "Wrong name")
			assert.Equal(t, render(cc.path, ctx), path, "Wrong path")
		})
	}
}
