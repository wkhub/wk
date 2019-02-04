package user

import (
	"path/filepath"

	"github.com/wkhub/wk/fs"
	"github.com/wkhub/wk/utils/config"
)

var (
	cfgFile string
	cfg     config.RawConfig
)

const CONFIG_FILENAME = "config.toml"

// Config is the
type Config struct {
	// *viper.Viper
}

func getUserConfig() config.RawConfig {
	if cfg == nil {
		cfg = config.New()
		cfg.AutomaticEnv() // read in environment variables that match
		cfg.SetEnvPrefix("wk")

		cfg.AddConfigPath(fs.Home())
		cfg.SetConfigName("config")
	}
	return cfg
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
