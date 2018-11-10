package fs

import (
	"os"
	"os/user"
	"path/filepath"
)

// Posix
// const XDG_CONFIG_HOME = %LOCALAPPDATA%
// const XDG_DATA_HOME = %LOCALAPPDATA%

// Windows
// const XDG_DATA_HOME = %LOCALAPPDATA%
// $XDG_DATA_DIRS = %APPDATA%
// const XDG_CONFIG_HOME = %LOCALAPPDATA%
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

// DataHome expose the user data root directory
func DataHome() string {
	return ""
}

func DataFilename(filename string) string {
	return ""
}

// ConfigHome expose the user config root directory
func ConfigHome() string {
	return ""
}

func ConfigFilename(filename string) string {
	return ""
}
