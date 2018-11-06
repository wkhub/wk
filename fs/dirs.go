package fs

import (
	"os"
	"os/user"
	"path/filepath"
)

// Windows
// $XDG_DATA_HOME = %LOCALAPPDATA%
// $XDG_DATA_DIRS = %APPDATA%
// $XDG_CONFIG_HOME = %LOCALAPPDATA%
// $XDG_CONFIG_DIRS = %APPDATA%
// $XDG_CACHE_HOME = %TEMP%
// $XDG_RUNTIME_DIR = %TEMP%

// Mac
// XDG_CONFIG_HOME ▶︎ ~/Library/Preferences/
// XDG_DATA_HOME ▶︎ ~/Library/
// XDG_CACHE_HOME ▶︎ ~/Library/Caches/
// Mapping XDG Base Directory Specification locations for "My App" on Mac OS X could look like this:

// XDG_CONFIG_HOME ▶︎ ~/Library/Preferences/name.often.with.domain.myapp.plist
// XDG_DATA_HOME ▶︎ ~/Library/My App/
// XDG_CACHE_HOME ▶︎ ~/Library/Caches/My App/

func Home() string {
	path := os.Getenv("WK_HOME")
	if path == "" {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(user.HomeDir, ".config", "wk")
	}
	return path
}

func ConfigFile(name string) string {
	return ""
}

func XdgConfigFile(name string) string {
	return ""
}

func WinConfigFile(name string) string {
	return ""
}

func MacConfigFile(name string) string {
	return ""
}

func DataFile(name string) string {
	return ""
}

func XdgDataFile(name string) string {
	return ""
}

func WinDataFile(name string) string {
	return ""
}

func MacDataFile(name string) string {
	return ""
}
