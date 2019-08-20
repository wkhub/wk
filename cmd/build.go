package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/user"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build project",
	Long:  "Shortcut for wk run build if defined",
	Run: func(cmd *cobra.Command, args []string) {
		user := user.Current()
		project := user.CurrentProject()

		if project == nil {
			log.Fatalf("No active project")
		}
		project.Load()
		key := "scripts.build"
		if !project.Config.IsSet(key) {
			log.Fatalf("No build script defined")
		}
		script := project.Config.GetString(key)

		session := shell.NewSession(false)
		session.AddCommand(script)
		code := user.Shell().Exec(session)
		os.Exit(code)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
