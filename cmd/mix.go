package cmd

import (
	"log"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/mixer"
)

// mixCmd represents the mix command
var mixCmd = &cobra.Command{
	Use:   "mix <source> [<target>]",
	Short: "Inject a wk template",
	Long:  `Open a subshell on the project path`,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		source, target := parseMixArgs(args)
		mixer := mixer.New(source)
		err := mixer.Mix(target)
		if err != nil {
			log.Fatal(err)
			// fmt.Println("Unkown project", name)
			os.Exit(1)
		}

		// target := os.Getcwd()
		// if len(args) == 2 {
		// 	name = args[1]
		// }

		// h := home.Get()
		// project := h.FindProject(name)
		// if project == nil {
		// 	fmt.Println("Unkown project", name)
		// 	os.Exit(1)
		// }
		// project.Open()
	},
}

func parseMixArgs(args []string) (string, string) {
	source := args[0]
	target, err := syscall.Getwd() // See: https://github.com/golang/go/issues/20947
	if err != nil {
		log.Fatal(err)
	}
	if len(args) == 2 {
		if filepath.IsAbs(args[1]) {
			target = args[1]
		} else {
			target = filepath.Join(target, args[1])
		}
	}
	return source, target
}

func init() {
	rootCmd.AddCommand(mixCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mixCmd.Flags().BoolP("force", "f", false, "Force overwrite")
}
