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

// mixCmd represents the mix command
var mixCmd = &cobra.Command{
	Use:   "mix <template>",
	Short: "Inject a wk template",
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
	rootCmd.AddCommand(mixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mixCmd.Flags().BoolP("force", "f", false, "Force overwrite")
}
