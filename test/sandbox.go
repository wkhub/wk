package test

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type SCtx struct {
	Cwd     string // Current working directoy
	Dirname string // dirname of the current working directory
}

func (ctx SCtx) Render(txt string) string {
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Sandbox creates a sandbox for testing wk environement
func Sandbox(cb func(ctx SCtx)) {
	cwd, err := syscall.Getwd() // See: https://github.com/golang/go/issues/20947
	handleError(err)
	tmp, err := ioutil.TempDir("", "")
	handleError(err)
	os.Chdir(tmp)
	ctx := SCtx{tmp, filepath.Base(tmp)}

	defer func() {
		os.Chdir(cwd)
	}()

	cb(ctx)
}
