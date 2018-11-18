package cmd

import (
	"fmt"
	"os"

	"github.com/wkhub/wk/shell"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/home"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on <project>",
	Short: "Work on a project",
	Long:  `Open a subshell on the project path`,
	Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		h := home.Get()
		project := h.FindProject(name)
		if project == nil {
			fmt.Println("Unknown project", name)
			os.Exit(1)
		}
		session := shell.NewSession(isEval)
		project.Contribute(&session)
		if isEval {
			shell.Current().Eval(session)
		} else {
			fmt.Printf("Opening project %s (%s)\n", project.Name, project.Root())
			shell.Current().Run(session)
			fmt.Printf("Exiting project %s\n", project.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.Flags().BoolP("ide", "i", false, "Launch ide")
}
