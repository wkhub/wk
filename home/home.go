package home

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/projects"
)

const _TEMPLATES_DIR = "templates"
const _SHELLS_DIR = "shells"

// Represent the WK_HOME directory
type Home struct {
	Path string
}

func Get() Home {
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
	return filepath.Join(h.Path, projects.BASE_DIR)
}

func (h Home) TemplatesDir() string {
	return filepath.Join(h.Path, _TEMPLATES_DIR)
}

func (h Home) ShellsDir() string {
	return filepath.Join(h.Path, _SHELLS_DIR)
}

func (h Home) Projects() []projects.Project {
	files, err := ioutil.ReadDir(h.ProjectsDir())
	if err != nil {
		log.Fatal(err)
	}
	var projs []projects.Project

	for _, file := range files {
		project := projects.NewFromFile(file.Name())
		projs = append(projs, project)
	}

	return projs
}

func (h Home) FindProject(name string) *projects.Project {
	for _, p := range h.Projects() {
		if p.Name == name {
			return &p
		}
	}
	return nil
}
