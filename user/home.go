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

func (h Home) Init() {
	if _, err := os.Stat(h.Path); os.IsNotExist(err) {
		fmt.Println("Creating WK_HOME in", h.Path)
		err = os.MkdirAll(h.Path, 0755)
		if err != nil {
			panic(err)
		}
	}
	os.MkdirAll(h.ProjectsDir(), 0755)
	os.MkdirAll(h.TemplatesDir(), 0755)
	os.MkdirAll(h.ShellsDir(), 0755)
}
