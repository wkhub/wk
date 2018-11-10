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

	"github.com/spf13/cobra"
)

// shrcCmd represents the completion command
var shrcCmd = &cobra.Command{
	Use:   "shrc",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(wk shrc --completion --aliases)

To configure your shell to load completions and aliases for each session add to your bashrc

# ~/.bashrc or ~/.zshrc or ~/.profile
. <(wk shrc --completion --aliases)

You can force the shell used by wk:
# Force bash
. <(wk shrc --bash --completion --aliases)
# Force zsh
. <(wk shrc --zsh --completion --aliases)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shrc")
		fmt.Println("com", cmd.Flag("completion").Value, &cmd.Flag("completion").Value)
		completion := cmd.Flag("completion").Value
		// aliases := cmd.Flag("aliases").Value
		// bash := cmd.Flag("bash").Value
		// zsh := cmd.Flag("zsh").Value
		// completion, err := cmd.Flags().GetBool("completion")
		// if err != nil {
		// 	fmt.Println("completion is nil")
		// }
		fmt.Sprintf("completion %s", completion)
		// aliases, err := cmd.Flags().GetBool("aliases")
		// fmt.Sprintf("aliases %s (%s)", aliases, err)
		// bash, err := cmd.Flags().GetBool("bash")
		// fmt.Sprintf("bash %s (%s)", bash, err)
		// zsh, err := cmd.Flags().GetBool("zsh")
		// fmt.Sprintf("zsh %s (%s)", zsh, err)
		// if (*bool)(completion) {
		// 	if bash {
		// 		rootCmd.GenBashCompletion(os.Stdout)
		// 	} else if zsh {
		// 		rootCmd.GenZshCompletion(os.Stdout)
		// 	} else {
		// 		shell.Current()
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(shrcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shrcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	shrcCmd.Flags().BoolP("bash", "", false, "Force bash syntax")
	shrcCmd.Flags().BoolP("zsh", "", false, "Force zsh syntax")

	shrcCmd.Flags().BoolP("completion", "", false, "Enable completion")
	shrcCmd.Flags().BoolP("aliases", "", false, "Enable aliases")
}
