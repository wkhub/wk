package config

import (
	"fmt"
	"io"
	"reflect"

	"github.com/spf13/viper"

	"github.com/wkhub/wk/utils/types"
)

type (
	// Config is the common config interface
	Config interface {
		// Get can retrieve any value given the key to use.
		// Get is case-insensitive for a key.
		// Get has the behavior of returning the value associated with the first
		// place from where it is set. Viper will check in the following order:
		// override, flag, env, config files, key/value store, default
		//
		// Get returns an interface. For a specific value use one of the Get____ methods.
		Get(key string) interface{}
		// Sub returns new Viper instance representing a sub tree of this instance.
		// Sub is case-insensitive for a key.
		Sub(key string) Config
		// GetString returns the value associated with the key as a string.
		GetString(key string) string
		// GetBool returns the value associated with the key as a boolean.
		GetBool(key string) bool
		// GetInt returns the value associated with the key as an integer.
		GetInt(key string) int
		// GetInt32 returns the value associated with the key as an integer.
		GetInt32(key string) int32
		// GetInt64 returns the value associated with the key as an integer.
		GetInt64(key string) int64
		// // GetFloat64 returns the value associated with the key as a float64.
		// GetFloat64(key string) float64
		// // GetTime returns the value associated with the key as time.
		// GetTime(key string) time.Time
		// // GetDuration returns the value associated with the key as a duration.
		// GetDuration(key string) time.Duration
		// GetStringSlice returns the value associated with the key as a slice of strings.
		GetStringSlice(key string) []string
		// GetMergedStringMap returns the value associated with the key as a map of interfaces.
		// Keys are merged from multiple sub configs
		GetMergedStringMap(key string) map[string]interface{}
		// GetStringMap returns the value associated with the key as a map of interfaces.
		GetStringMap(key string) map[string]interface{}
		// GetStringMapString returns the value associated with the key as a map of strings.
		GetStringMapString(key string) map[string]string
		// GetMergedStringMapString returns the value associated with the key as a map of strings.
		// Keys are merged from multiple sub configs
		GetMergedStringMapString(key string) map[string]string
		// // GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
		// GetStringMapStringSlice(key string) map[string][]string
		// // GetSizeInBytes returns the size of the value associated with the given key
		// // in bytes.
		// GetSizeInBytes(key string) uint
		// // UnmarshalKey takes a single key and unmarshals it into a Struct.
		// UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
		// Unmarshal unmarshals the config into a Struct. Make sure that the tags
		// on the fields of the structure are properly set.
		Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
		// Set sets the value for the key in the override register.
		// Set is case-insensitive for a key.
		// Will be used instead of values obtained via
		// flags, config file, ENV, default, or key/value store.
		Set(key string, value interface{})
		// IsSet checks to see if the key has been set in any of the data locations.
		// IsSet is case-insensitive for a key.
		IsSet(key string) bool
		// // InConfig checks to see if the given key (or an alias) is in the config file.
		// InConfig(key string) bool
		// // SetDefault sets the default value for this key.
		// // SetDefault is case-insensitive for a key.
		// // Default only used when no value is provided by the user via flag, config or ENV.
		// SetDefault(key string, value interface{})
		// AllKeys returns all keys holding a value, regardless of where they are set.
		// Nested keys are returned with a v.keyDelim (= ".") separator
		AllKeys() []string
		// AllSettings merges all settings and returns them as a map[string]interface{}.
		AllSettings() map[string]interface{}
		// Load loads the configuration file if exists
		Load()
	}

	// RawConfig represent a classic *viper.Viper
	RawConfig interface {
		Config
		SetConfigName(in string)
		SetConfigType(in string)
		SetConfigFile(in string)
		AddConfigPath(in string)
		ConfigFileUsed() string
		AutomaticEnv()
		SetEnvPrefix(in string)
		ReadConfig(in io.Reader) error
		WriteConfig() error
		WriteConfigAs(filename string) error
	}

	cascade struct {
		configs []Config
	}

	config struct {
		*viper.Viper
	}
)

func (cfg config) Sub(key string) Config {
	sub := cfg.Viper.Sub(key)
	if sub != nil {
		return config{sub}
	}
	return nil
}

// New returns a new simple Config instance
func New() RawConfig {
	return config{viper.New()}
}

func (cfg config) Load() {
	if err := cfg.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// Allow missing config
			return
		default:
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	// return cfg.ConfigFileUsed()
}

func (cfg config) GetMergedStringMap(key string) map[string]interface{} {
	return cfg.GetStringMap(key)
}

func (cfg config) GetMergedStringMapString(key string) map[string]string {
	return cfg.GetStringMapString(key)
}

// Cascade returns a new cascading Config instance
func Cascade(configs ...Config) Config {
	return cascade{configs}
}

func callByName(cfg Config, method string, key string) interface{} {
	meth := reflect.ValueOf(cfg).MethodByName(method)
	params := []reflect.Value{reflect.ValueOf(key)}
	return meth.Call(params)[0].Interface()
}

func (cfg cascade) first(method string, key string) interface{} {
	var last Config
	for _, config := range cfg.configs {
		if config.IsSet(key) {
			return callByName(config, method, key)
		}
		last = config
	}
	return callByName(last, method, key)
}

func (cfg cascade) Get(key string) interface{} {
	return cfg.first("Get", key)
}

func (cfg cascade) Set(key string, value interface{}) {
	cfg.configs[0].Set(key, value)
}

func (cfg cascade) Sub(key string) Config {
	var subs []Config
	for _, config := range cfg.configs {
		sub := config.Sub(key)
		if sub != nil {
			subs = append(subs, sub)
		}
	}
	return Cascade(subs...)
}

func (cfg cascade) GetString(key string) string {
	return cfg.first("GetString", key).(string)
}

func (cfg cascade) GetBool(key string) bool {
	return cfg.first("GetBool", key).(bool)
}

func (cfg cascade) GetInt(key string) int {
	return cfg.first("GetInt", key).(int)
}

func (cfg cascade) GetInt32(key string) int32 {
	return cfg.first("GetInt32", key).(int32)
}

func (cfg cascade) GetInt64(key string) int64 {
	return cfg.first("GetInt64", key).(int64)
}

func (cfg cascade) GetStringSlice(key string) []string {
	return cfg.first("GetStringSlice", key).([]string)
}

func (cfg cascade) GetStringMap(key string) map[string]interface{} {
	return cfg.first("GetStringMap", key).(map[string]interface{})
}

func (cfg cascade) GetMergedStringMap(key string) map[string]interface{} {
	result := map[string]interface{}{}
	for idx := range cfg.configs {
		config := cfg.configs[len(cfg.configs)-1-idx]
		for key, value := range config.GetStringMap(key) {
			result[key] = value
		}
	}
	return result
}

func (cfg cascade) GetStringMapString(key string) map[string]string {
	return cfg.first("GetStringMapString", key).(map[string]string)
}

func (cfg cascade) GetMergedStringMapString(key string) map[string]string {
	result := map[string]string{}
	for idx := range cfg.configs {
		config := cfg.configs[len(cfg.configs)-1-idx]
		for key, value := range config.GetStringMapString(key) {
			result[key] = value
		}
	}
	return result
}

func (cfg cascade) IsSet(key string) bool {
	for _, config := range cfg.configs {
		if config.IsSet(key) {
			return true
		}
	}
	return false
}

func (cfg cascade) Load() {}

func (cfg cascade) reversedConfigs() []Config {
	out := []Config{}
	for i := range cfg.configs {
		out = append(out, cfg.configs[len(cfg.configs)-i-1])
	}
	return out
}

func (cfg cascade) Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
	for _, config := range cfg.reversedConfigs() {
		err := config.Unmarshal(rawVal)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cfg cascade) AllSettings() map[string]interface{} {
	out := map[string]interface{}{}
	for _, key := range cfg.AllKeys() {
		out[key] = cfg.Get(key)
	}
	return out
}

func (cfg cascade) AllKeys() []string {
	out := []string{}
	keys := types.Set()
	for _, config := range cfg.configs {
		for _, key := range config.AllKeys() {
			keys.Add(key)
		}
	}

	for _, key := range keys.Slice() {
		out = append(out, key.(string))
	}
	return out
}
