package mixer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar"
	"github.com/pkg/errors"

	"github.com/wkhub/wk/mixer/backends"
	"github.com/wkhub/wk/utils/config"
)

type Mixer struct {
	Path string
}

func (m Mixer) TemplateRoot() string {
	return filepath.Join(m.Path, "template")
}

// Mix reads the mixer, prompt the parameters to user
//	and then apply it the target
func (m Mixer) Mix(target string) error {
	// cfgFilename := filepath.Join(m.Path, "mixer.toml")
	cfg := config.New()
	cfg.AddConfigPath(m.Path)
	cfg.SetConfigName("mixer")
	// cfg.SetConfigFile(cfgFilename)
	// cfg.SetConfigType("toml")
	cfg.Load()
	// err := cfg.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }
	config := Config{}
	err := cfg.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal unmarshalling config: %s", err))
	}

	ctx := Context{}

	config.Params.PromptUser(ctx)

	tplRoot := m.TemplateRoot()

	ignoreList, err := ctx.RenderList(config.Mix.Ignore)
	if err != nil {
		return errors.Wrapf(err, `Unable to render ignore list '%s'`, config.Mix.Ignore)
	}
	copyList, err := ctx.RenderList(config.Mix.Copy)
	if err != nil {
		return errors.Wrapf(err, `Unable to render copy list '%s'`, config.Mix.Copy)
	}

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
			return errors.Wrapf(err, `Unable to compute path '%s'`, path)
		}
		relTargetPath, err := ctx.Render(relPath)
		if err != nil {
			fmt.Printf("relTarget: %v\n", err)
			return errors.Wrapf(err, `Unable to compute target path '%s'`, relPath)
		}

		if relTargetPath == "" || Match(relTargetPath, ignoreList) {
			// Skip empty paths and ignore list
			return filepath.SkipDir
		}

		targetPath := filepath.Join(target, relTargetPath)

		if info.Mode().IsRegular() {
			fmt.Println("Processing", relPath)
			if err = os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return err
			}
			if Match(relTargetPath, copyList) {
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				if err = ioutil.WriteFile(targetPath, bytes, info.Mode()); err != nil {
					return err
				}
			} else {
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Fatal(err)
				}

				content, err := ctx.Render(string(bytes))
				if err != nil {
					fmt.Printf("content: %v\n", err)
					return err
				}

				if err = ioutil.WriteFile(targetPath, []byte(content), info.Mode()); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", m.TemplateRoot(), err)
		return err
	}

	// TODO:
	//	 - Postmix scripts/hooks
	return nil
}

func New(source string) Mixer {
	backend, err := backends.Resolve(source)
	if err != nil {
		panic(err)
	}
	fmt.Println("Using backend", backend.Name())
	fetched, err := backend.Fetch(source)
	if err != nil {
		panic(err)
	}
	return Mixer{fetched}
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
