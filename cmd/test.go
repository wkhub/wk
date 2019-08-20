package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/user"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute tests",
	Long:  "Shortcut for wk run test if defined",
	Run: func(cmd *cobra.Command, args []string) {
		user := user.Current()
		project := user.CurrentProject()

		if project == nil {
			log.Fatalf("No active project")
		}
		project.Load()
		key := "scripts.test"
		if !project.Config.IsSet(key) {
			log.Fatalf("No test script defined in wk.yaml")
		}
		script := project.Config.GetString(key)

		session := shell.NewSession(false)
		session.AddCommand(script)
		code := user.Shell().Exec(session)
		os.Exit(code)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
