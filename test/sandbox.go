package test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type SCtx struct {
	Cwd     string // Current working directoy
	Dirname string // dirname of the current working directory
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Sandbox creates a sandbox for testing wk environement
func Sandbox(cb func(ctx SCtx)) {
	cwd, err := os.Getwd()
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
