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
	"os"

	"github.com/noirbizarre/wk/home"
	"github.com/spf13/cobra"
)

// rmCmd represents the new command
var rmCmd = &cobra.Command{
	Use:   "rm <project>",
	Short: "Remove a new project",
	Long: `Remove a project definition.

	Project root is left untouched
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		h := home.Get()
		project := h.FindProject(name)
		if project == nil {
			fmt.Println("Unkown project", name)
			os.Exit(1)
		}
		project.Delete()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// rmCmd.Flags().String("mix", "m", "Mix a template")
}
