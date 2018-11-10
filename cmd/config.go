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

	"github.com/wkhub/wk/home"
	"github.com/spf13/cobra"
)

// configCmd represents the set command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Associate a project with the current directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

// initCmd represents the from command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a configuration file",
	Long:  `Initialize a wk configuration either globaly, bla bla`,
	Run: func(cmd *cobra.Command, args []string) {
		h := home.Get()
		for _, project := range h.Projects() {
			fmt.Println(project.Name)
		}
	},
}

// getCmd represents the set command
var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration key",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration key",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("set called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	configCmd.PersistentFlags().BoolP("global", "g", false, "Set a global config key")
	configCmd.PersistentFlags().BoolP("local", "l", false, "Set a local project config key")
	configCmd.PersistentFlags().BoolP("shared", "s", false, "Set a shared project config key")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
