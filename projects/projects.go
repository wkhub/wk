package projects

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/utils/config"
)

const BASE_DIR = "projects"

// Project metadata
type Project struct {
	Name         string
	UserConfig   config.RawConfig
	SharedConfig config.RawConfig
	LocalConfig  config.RawConfig
	Config       config.Config
}

func (p *Project) ensureConfig() {
	if p.Config == nil {
		p.UserConfig = config.New()
		p.UserConfig.AddConfigPath(filepath.Join(fs.Home(), BASE_DIR))
		p.UserConfig.SetConfigName(p.Name)

		p.SharedConfig = config.New()

		p.LocalConfig = config.New()

		p.Config = config.Cascade(p.UserConfig, p.SharedConfig, p.LocalConfig)
	}
}

// Root resturns the project root path
func (p Project) Root() string {
	path := p.UserConfig.GetString("path")
	if filepath.IsAbs(path) {
		return path
	} else {
		return filepath.Join(fs.Projects(), path)
	}
}

// Load loads project configuration from file
func (p *Project) Load() {
	p.ensureConfig()
	p.UserConfig.Load()

	p.SharedConfig.AddConfigPath(p.Root())
	p.SharedConfig.SetConfigName("wk")
	p.SharedConfig.Load()

	p.LocalConfig.AddConfigPath(p.Root())
	p.LocalConfig.SetConfigName("wk.local")
	p.LocalConfig.Load()
}

// Contribute enrich a session for a given project
func (p Project) Contribute(session *shell.Session) *shell.Session {
	p.Load()
	session.Cwd = p.Root()
	session.Env["WK_PROJECT"] = p.Name
	for key, value := range p.Config.GetMergedStringMapString("env") {
		session.Env[strings.ToUpper(key)] = value
	}
	// if !p.Config.IsSet("env") {
	// 	// for _, line := range strings.Split(string(text), "\n") {
	// 	// 	if strings.TrimSpace(line) != "" {
	// 	// 		parts := strings.Split(line, "=")
	// 	// 		session.Env[parts[0]] = parts[1]
	// 	// 	}
	// 	// }
	// }

	// return hooks.Execute(session)
	return session
}

// Save create a project or persists its changes
func (p *Project) Save() {
	p.ensureConfig()
	fmt.Println(p.UserConfig.ConfigFileUsed())
	p.UserConfig.WriteConfig()
}

// Delete remove a project definition
func (p Project) Delete() {
	p.ensureConfig()
	file := p.UserConfig.ConfigFileUsed()
	fmt.Printf("Deleting project %s (%s)\n", p.Name, file)
	err := os.Remove(file)
	if err != nil {
		panic(err)
	}
}

// New initialize a project
func New(name string) Project {
	project := Project{Name: name}
	project.ensureConfig()
	return project
}

// NewFromFile initialize a Project from a path
func NewFromFile(pth string) Project {
	name := strings.TrimSuffix(pth, path.Ext(pth))
	return New(name)
}

// Create instantiate a new project
func Create(name string, path string) Project {
	cfg := config.New()
	cfg.Set("path", path)
	cfg.SetConfigType("toml")
	filename := fmt.Sprintf("%s.toml", name)
	filename = filepath.Join(fs.Home(), BASE_DIR, filename)
	if err := cfg.WriteConfigAs(filename); err != nil {
		panic(err)
	}

	if !filepath.IsAbs(path) {
		path = filepath.Join(fs.Projects(), path)
	}

	if !fs.Exists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	return New(name)
}
