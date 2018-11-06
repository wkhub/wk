package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/noirbizarre/wk/fs"
)

// Represent the WK_HOME directory
type Home struct {
	Path string
}

func GetHome() Home {
	homePath := fs.Home()
	home := Home{homePath}
	home.Init()
	return home
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

func (h Home) Config() string {
	return filepath.Join(h.Path, CONFIG_FILENAME)
}

func (h Home) ProjectsDir() string {
	return filepath.Join(h.Path, "projects")
}

func (h Home) TemplatesDir() string {
	return filepath.Join(h.Path, "templates")
}

func (h Home) ShellsDir() string {
	return filepath.Join(h.Path, "shells")
}

func (h Home) Projects() []Project {
	files, err := ioutil.ReadDir(h.ProjectsDir())
	if err != nil {
		log.Fatal(err)
	}
	var projects []Project

	for _, file := range files {
		project := ProjectFromFile(file.Name())
		projects = append(projects, project)
	}

	return projects
}

func (h Home) FindProject(name string) *Project {
	for _, p := range h.Projects() {
		if p.Name == name {
			return &p
		}
	}
	return nil
}
