package user

import (
	"fmt"
	"os"

	"github.com/wkhub/wk/fs"
)

// WkHome represent the WK_HOME directory
type Home struct {
	Path string
}

func WkHome() Home {
	path := fs.WkHome()
	wkhome := Home{path}
	wkhome.Init()
	return wkhome
}

func mkdir(path string, perm int) {
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}
}

func (h Home) Init() {
	if _, err := os.Stat(h.Path); os.IsNotExist(err) {
		fmt.Println("Creating WK_HOME in", h.Path)
		mkdir(h.Path, 0755)
	}
	mkdir(h.ProjectsDir(), 0755)
	mkdir(h.TemplatesDir(), 0755)
	mkdir(h.ShellsDir(), 0755)
}
