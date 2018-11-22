package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config interface {
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
	Sub(key string) *Config
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
	// GetFloat64 returns the value associated with the key as a float64.
	GetFloat64(key string) float64
	// GetTime returns the value associated with the key as time.
	GetTime(key string) time.Time
	// GetDuration returns the value associated with the key as a duration.
	GetDuration(key string) time.Duration
	// GetStringSlice returns the value associated with the key as a slice of strings.
	GetStringSlice(key string) []string
	// GetStringMap returns the value associated with the key as a map of interfaces.
	GetStringMap(key string) map[string]interface{}
	// GetStringMapString returns the value associated with the key as a map of strings.
	GetStringMapString(key string) map[string]string
	// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
	GetStringMapStringSlice(key string) map[string][]string
	// GetSizeInBytes returns the size of the value associated with the given key
	// in bytes.
	GetSizeInBytes(key string) uint
	// UnmarshalKey takes a single key and unmarshals it into a Struct.
	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
	// Unmarshal unmarshals the config into a Struct. Make sure that the tags
	// on the fields of the structure are properly set.
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	// IsSet checks to see if the key has been set in any of the data locations.
	// IsSet is case-insensitive for a key.
	IsSet(key string) bool
	// InConfig checks to see if the given key (or an alias) is in the config file.
	InConfig(key string) bool
	// SetDefault sets the default value for this key.
	// SetDefault is case-insensitive for a key.
	// Default only used when no value is provided by the user via flag, config or ENV.
	SetDefault(key string, value interface{})
	// AllKeys returns all keys holding a value, regardless of where they are set.
	// Nested keys are returned with a v.keyDelim (= ".") separator
	AllKeys() []string
	// AllSettings merges all settings and returns them as a map[string]interface{}.
	AllSettings() map[string]interface{}
}

type CascadingConfig struct {
	configs []Config
}

func (cfg CascadingConfig) Get(key string) interface{} {
	for _, config := range cfg.configs {
		if config.IsSet(key) {
			return config.Get(key)
		}
	}
	return nil
}

func (cfg CascadingConfig) IsSet(key string) bool {
	for _, config := range cfg.configs {
		if config.IsSet(key) {
			return true
		}
	}
	return false
}

// func (cfg CascadingConfig) Sub()
