package user

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wkhub/wk/projects"
)

var (
	currentProject string
)

func (h Home) ProjectsDir() string {
	return filepath.Join(h.Path, projects.BASE_DIR)
}

func (user User) Projects() []projects.Project {
	files, err := ioutil.ReadDir(user.Home.ProjectsDir())
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

// FindProject loops through user project and return the first matching name if any
func (user User) FindProject(name string) *projects.Project {
	for _, p := range user.Projects() {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func (user User) SetProject(name string) *projects.Project {
	project := user.FindProject(name)
	if project == nil {
		fmt.Println("Unknown project", name)
	} else {
		currentProject = name
	}
	return project
}

func (user User) CurrentProject() *projects.Project {
	if currentProject == "" {
		currentProject = os.Getenv("WK_PROJECT")
	}
	return user.FindProject(currentProject)
}
