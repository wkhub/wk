package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/user"
)

// cdCmd represents the set command
var cdCmd = &cobra.Command{
	Use:   "cd <dir>",
	Short: "Go to a predefined directory",
	Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		user := user.Current()
		project := user.CurrentProject()

		target := args[0]

		switch target {
		case "home":
			path = user.Home.Path
		case "projects":
			path = user.Home.ProjectsDir()
		case "project":
			if project != nil {
				project.Load()
				path = project.Root()
			}
		default:
			user.Config.Load()
			key := fmt.Sprintf("dirs.%s", target)
			if user.Config.IsSet(key) {
				path = user.Config.GetString(key)
			}
		}
		if path == "" {
			fmt.Println("Unknown target", target)
			return
		}
		session := shell.NewSession(isEval)
		initCmd := fmt.Sprintf("cd %s", path)
		session.Init = append(session.Init, initCmd)
		if isEval {
			user.Shell().Eval(session)
		} else {
			user.Shell().Run(session)
		}

	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
