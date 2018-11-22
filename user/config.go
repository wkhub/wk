package user

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/wkhub/wk/fs"
)

var (
	cfgFile string
	config  *UserConfig
)

const CONFIG_FILENAME = "config.toml"

// Config is the
type Config struct {
	// *viper.Viper
}

type UserConfig struct {
	*viper.Viper
}

func getUserConfig() *UserConfig {
	if config == nil {
		config = &UserConfig{viper.New()}
		config.AutomaticEnv() // read in environment variables that match
		config.SetEnvPrefix("wk")

		config.AddConfigPath(fs.Home())
		config.SetConfigName("config")
		config.SetConfigType("toml")
	}
	return config
}

// Load loads project configuration from file
func (config UserConfig) Load() {
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func (h Home) ConfigPath() string {
	return filepath.Join(h.Path, CONFIG_FILENAME)
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	// home := pkg.GetHome()
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		h := home.Get()
// 		viper.AddConfigPath(h.Path)
// 		viper.SetConfigName(home.CONFIG_FILENAME)
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match
// 	viper.SetEnvPrefix("wk")

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
