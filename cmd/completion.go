package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates completion scripts",
	Long: `To load completion run

On Bash, add this to your .bashrc
. <(wk completion --bash)

On Zsh, you need to create a completion file.
Given $FPATH a path in $fpath:
wk completion --zsh > $FPATH/_wk
`,
	Run: func(cmd *cobra.Command, args []string) {
		checkShellFlags()
		if isZsh {
			rootCmd.GenZshCompletion(os.Stdout)
		} else if isBash {
			rootCmd.GenBashCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
