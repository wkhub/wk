package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/shell"
)

var (
	shrcCompletion bool
	shrcAliases    bool
)

// shrcCmd represents the completion command
var shrcCmd = &cobra.Command{
	Use:   "shrc",
	Short: "Initialize wk shell integration",
	Long: `Load wk in the shell allowing to not spawn subshells.

On bash, in your ~/.bashrc
. <(wk shrc --bash)

If you also want to export completion:
. <(wk shrc --bash --completion)

On zsh, in your ~/.zshrc
. <(wk shrc --zsh)
`,
	Run: func(cmd *cobra.Command, args []string) {
		checkShellFlags()
		if shrcCompletion && !isBash {
			fmt.Println("Only bash support sourcing completion")
			os.Exit(1)
		}
		// Tells wk than shell has been initialized
		fmt.Println("export WK_IN_SHELL=true")
		if isZsh {
			fmt.Println(shell.ZSH.Rc())
		} else if isBash {
			if shrcCompletion {
				if err := rootCmd.GenBashCompletion(os.Stdout); err != nil {
					panic(err)
				}
			}
			fmt.Println(shell.BASH.Rc())
		}
	},
}

func init() {
	rootCmd.AddCommand(shrcCmd)

	shrcCmd.Flags().BoolVar(&shrcCompletion, "completion", false, "Enable completion")
	shrcCmd.Flags().BoolVar(&shrcAliases, "aliases", false, "Enable aliases")
}
