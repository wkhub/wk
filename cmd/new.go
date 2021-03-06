package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/projects"
	"github.com/wkhub/wk/shell"
	"github.com/wkhub/wk/user"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Long: `Create project from scratch or from content.

	By default the project root will be created into WK_PROJECT
	You can override this behavior with the --target option.

	# create a new project named my-project in $WK_PROJECTS:
	wk new my-project

	# create a new project named my-project in the current directory:
	wk new my-project .

	# create a new project named in the current directory, name is guess from directory:
	wk new .

	# create a new project named in a given directory, name is guess from directory:
	wk new path/to/project
	`,
	Args: cobra.RangeArgs(1, 2),
	Annotations: map[string]string{
		"source": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		name, path := newGuessArgs(args)
		project := projects.Create(name, path)
		session := shell.NewSession(isEval)
		project.Contribute(&session)
		if cmd.Flag("mix").Changed {
			mixCmd := fmt.Sprintf("wk mix %s", cmd.Flag("mix").Value.String())
			session.Init = append(session.Init, mixCmd)
		}
		if isEval {
			user.Current().Shell().Eval(session)
		} else {
			fmt.Printf("Opening project %s (%s)\n", project.Name, project.Root())
			user.Current().Shell().Run(session)
			fmt.Printf("Exiting project %s\n", project.Name)
		}
	},
}

func newGuessArgs(args []string) (string, string) {
	var name string
	var path string
	var err error
	if len(args) == 2 {
		name = args[0]
		path, err = filepath.Abs(args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		arg := args[0]
		if fs.IsDir(arg) {
			// It's an existing directory
			path, err = filepath.Abs(arg)
			if err != nil {
				fmt.Printf("Unkown error %s\n", err)
			}
			name = filepath.Base(path)
		} else {
			// It's a project name
			name = arg
			path = arg
		}
	}
	return name, path
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("mix", "m", "", "Mix a template")
}
