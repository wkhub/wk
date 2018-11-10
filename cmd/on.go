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

	"github.com/spf13/cobra"

	"github.com/noirbizarre/wk/home"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on <project>",
	Short: "Work on a project",
	Long:  `Open a subshell on the project path`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		h := home.Get()
		project := h.FindProject(name)
		if project == nil {
			fmt.Println("Unkown project", name)
			os.Exit(1)
		}
		project.Open()
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.Flags().BoolP("ide", "i", false, "Launch ide")
}
