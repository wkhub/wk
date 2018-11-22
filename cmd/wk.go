package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wkhub/wk/user"
)

var (
	// VERSION is set during build
	VERSION        string
	cfgFile        string
	wkHome         string
	currentProject string
	verbose        bool
	isZsh          bool
	isBash         bool
	isEval         bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wk",
	Short: "Language agnostic project manager",
	Long: `wk allows to manage project:

Use wk on to switch on an existing project.
Use wk new to create a new project.
Use wk set to attach a project to the current directory`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default is $HOME/.config/wk/config.toml)")

	rootCmd.PersistentFlags().StringVar(&wkHome, "home", "",
		"config file (default is $HOME/.config/wk/)")

	// Shell
	rootCmd.PersistentFlags().BoolVar(&isBash, "bash", false, "Use bash syntax")
	rootCmd.PersistentFlags().BoolVar(&isZsh, "zsh", false, "Use zsh syntax")
	rootCmd.PersistentFlags().BoolVar(&isEval, "eval", false, "Return result to be called with eval")
	rootCmd.PersistentFlags().MarkHidden("eval")

	// Project
	rootCmd.PersistentFlags().StringVarP(&currentProject, "project", "p", "", "Set an active project")
}

func checkShellFlags() {
	if isBash && isZsh {
		fmt.Println("You can only specify one shell")
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	user := user.Current()
	if cfgFile != "" {
		user.Config.SetConfigFile(cfgFile)
	}
	// else if home
	// home := pkg.GetHome()
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	user.Config.SetConfigFile(cfgFile)
	// 	// viper.SetConfigFile(cfgFile)
	// } else {
	// 	h := home.Get()
	// 	viper.AddConfigPath(h.Path)
	// 	viper.SetConfigName(home.CONFIG_FILENAME)
	// }

	// viper.AutomaticEnv() // read in environment variables that match
	// viper.SetEnvPrefix("wk")

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
