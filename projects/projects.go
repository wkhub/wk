package projects

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/hooks"
	"github.com/wkhub/wk/shell"
)

const BASE_DIR = "projects"

// Project metadata
type Project struct {
	Name   string
	Config *viper.Viper
}

func (p *Project) ensureConfig() {
	if p.Config == nil {
		p.Config = viper.New()
		p.Config.SetConfigFile(p.Filename())
		p.Config.SetConfigType("toml")
	}
}

// Root resturns the project root path
func (p Project) Root() string {
	path := p.Config.GetString("path")
	if filepath.IsAbs(path) {
		return path
	} else {
		return filepath.Join(fs.Projects(), path)
	}
}

// Load loads project configuration from file
func (p *Project) Load() {
	p.ensureConfig()
	err := p.Config.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// Open opens a shell for a given project
func (p Project) Open() {
	p.Load()
	root := p.Root()
	fmt.Printf("Opening project %s (%s)\n", p.Name, root)
	env, sh := hooks.Execute(root)
	shell.Current().Run(root, env, sh)
	fmt.Printf("Exiting project %s\n", p.Name)
}

// Save create a project or persists its changes
func (p *Project) Save() {
	p.ensureConfig()
	p.Config.WriteConfig()
}

// Delete remove a project definition
func (p Project) Delete() {
	fmt.Printf("Deleting project %s (%s)\n", p.Name, p.Filename())
	err := os.Remove(p.Filename())
	if err != nil {
		panic(err)
	}
}

// Filename gives the project config filename
func (p Project) Filename() string {
	basename := fmt.Sprintf("%s.toml", p.Name)
	return filepath.Join(fs.Home(), BASE_DIR, basename)
}

// New initialize a project
func New(name string) Project {
	project := Project{name, nil}
	project.ensureConfig()
	return project
}

// NewFromFile initialize a Project from a path
func NewFromFile(pth string) Project {
	name := strings.TrimSuffix(pth, path.Ext(pth))
	project := Project{name, nil}
	project.ensureConfig()
	return project
}
