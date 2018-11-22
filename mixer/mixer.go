package mixer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/spf13/viper"
)

type Mixer struct {
	Path string
}

func (m Mixer) TemplateRoot() string {
	return filepath.Join(m.Path, "template")
}

func (m Mixer) Mix(target string) error {
	cfgFilename := filepath.Join(m.Path, "mixer.toml")
	cfg := viper.New()
	cfg.SetConfigFile(cfgFilename)
	cfg.SetConfigType("toml")
	err := cfg.ReadInConfig() // Find and read the config file
	if err != nil {           // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	config := Config{}
	err = cfg.Unmarshal(&config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal unmarshalling config: %s \n", err))
	}

	ctx := Context{}

	config.Params.PromptUser(ctx)

	tplRoot := m.TemplateRoot()

	err = filepath.Walk(tplRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if path == tplRoot {
			return nil
		}

		if info.IsDir() && info.Name() == "skip" {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		relPath, err := filepath.Rel(tplRoot, path)
		if err != nil {
			fmt.Printf("rel: %v\n", err)
			return err
		}
		targetPath := filepath.Join(target, ctx.Render(relPath))

		if Match(relPath, config.Mix.Ignore) {
			// Skip ignore list
			return filepath.SkipDir
		}

		if info.Mode().IsRegular() {
			fmt.Println("Processing", relPath)
			os.MkdirAll(filepath.Dir(targetPath), 0755)
			if Match(relPath, config.Mix.Copy) {
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				ioutil.WriteFile(targetPath, bytes, info.Mode())
			} else {
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				content := ctx.Render(string(bytes))

				ioutil.WriteFile(targetPath, []byte(content), info.Mode())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", m.TemplateRoot(), err)
		return err
	}
	return nil
}

func New(source string) Mixer {
	return Mixer{source}
}

func Match(path string, patterns []string) bool {
	for _, pattern := range patterns {
		match, _ := doublestar.PathMatch(pattern, path)
		if match {
			return true
		}
	}
	return false
}
