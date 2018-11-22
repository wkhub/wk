package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/user"
)

// linkCmd represents the set command
var linkCmd = &cobra.Command{
	Use:   "link <target>",
	Short: "Open/follow a predefined link",
	Args:  cobra.ExactArgs(1),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		var uri string
		target := args[0]
		key := fmt.Sprintf("uris.%s", target)
		user := user.Current()
		project := user.CurrentProject()

		if project != nil {
			project.Load()
			if project.Config.IsSet(key) {
				uri = project.Config.GetString(key)
			}
		}
		if uri == "" {
			user.Config.Load()
			if user.Config.IsSet(key) {
				uri = user.Config.GetString(key)
			}
		}
		if uri == "" {
			fmt.Println("Unknown target", target)
			return
		}
		if runtime.GOOS == "linux" {
			_, err := exec.Command("xdg-open", uri).Output()
			if err != nil {
				log.Fatal(err)
			}
		} else if runtime.GOOS == "darwin" {
			_, err := exec.Command("open", uri).Output()
			if err != nil {
				log.Fatal(err)
			}
		} else if runtime.GOOS == "windows" {
			_, err := exec.Command("start", uri).Output()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
