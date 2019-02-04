package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/user"
)

// runCmd represents the set command
var runCmd = &cobra.Command{
	Use:   "run <script>",
	Short: "Execute a predefined script",
	Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := user.Current()
		project := user.CurrentProject()

		if project == nil {
			log.Fatalf("No active project")
		}

		name := args[0]

		project.Load()
		key := fmt.Sprintf("scripts.%s", name)
		if !project.Config.IsSet(key) {
			log.Fatalf("Unknown script %s", name)
		}
		script := project.Config.GetString(key)

		session := shell.NewSession(false)
		session.AddCommand(script)
		code := user.Shell().Exec(session)
		os.Exit(code)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
