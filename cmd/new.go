// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/projects"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Long: `Create project from scratch or from content.

	By default the project root will be created into WK_PROJECT
	You can override this behavior with the --target option.

	# create a new project named my-project in $WK_PROJECTS:
	wk new my-project

	# create a new project named my-project in the current directory:
	wk new my-project .

	# create a new project named in the current directory, name is guess from directory:
	wk new .

	# create a new project named in a given directory, name is guess from directory:
	wk new path/to/project
	`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		name, path := newGuessArgs(args)
		fmt.Printf("Should create project %s in %s\n", name, path)
		project := projects.New(name)
		project.Config.Set("path", path)
		project.Save()
		root := project.Root()
		if !fs.Exists(root) {
			os.MkdirAll(root, os.ModePerm)
		}
		project.Open()
	},
}

func newGuessArgs(args []string) (string, string) {
	var name string
	var path string
	var err error
	if len(args) == 2 {
		name = args[0]
		path, err = filepath.Abs(args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		arg := args[0]
		if fs.IsDir(arg) {
			// It's an existing directory
			path, err = filepath.Abs(arg)
			if err != nil {
				fmt.Printf("Unkown error %s\n", err)
			}
			name = filepath.Base(path)
		} else {
			// It's a project name
			name = arg
			path = arg
		}
	}
	return name, path
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().String("mix", "m", "Mix a template")
}
