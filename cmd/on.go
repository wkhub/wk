package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/home"
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
		if isEval {
			fmt.Println(project.OpenIn())
		} else {
			project.Open()
		}
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
	onCmd.Flags().BoolP("ide", "i", false, "Launch ide")
}
