package projects

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/viper"

	"github.com/noirbizarre/wk/hooks"
	"github.com/noirbizarre/wk/shell"
)

// Project metadata
type Project struct {
	Filename string
	Name     string
	Config   *viper.Viper
}

func (p *Project) Load() {
	home := GetHome()
	p.Config = viper.New()
	p.Config.AddConfigPath(home.ProjectsDir())
	p.Config.SetConfigName(p.Name)
	p.Config.SetConfigType("toml")
	err := p.Config.ReadInConfig() // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (p Project) Open() {
	p.Load()
	cwd := p.Config.GetString("path")
	fmt.Printf("Opening project %s (%s)\n", p.Name, p.Config.Get("path"))
	env, sh := hooks.Execute(cwd)
	// fmt.Println(">> Starting a new interactive shell")
	shell.Current().Run(cwd, env, sh)
	fmt.Printf("Exiting project %s\n", p.Name)
}

// New initialize a project
func New(name string) Project {
	filename := name + ".toml"
	// basename := strings.TrimSuffix(pth, path.Ext(pth))
	project := Project{filename, name, nil}
	return project
}

// NewFromFile initialize a Project from a path
func NewFromFile(pth string) Project {
	basename := strings.TrimSuffix(pth, path.Ext(pth))
	project := Project{pth, basename, nil}
	return project
}
