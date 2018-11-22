package fs

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func currentUser() *user.User {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user
}

/*
 * Linux standard dirs
 */

func XdgConfigHome(user *user.User) string {
	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME == "" {
		return filepath.Join(user.HomeDir, ".config")
	}
	return XDG_CONFIG_HOME
}

func XdgCacheHome(user *user.User) string {
	XDG_CACHE_HOME := os.Getenv("XDG_CACHE_HOME")
	if XDG_CACHE_HOME == "" {
		return filepath.Join(user.HomeDir, ".cache")
	}
	return XDG_CACHE_HOME
}

/*
* MacOS standard dirs
 */
// Mac
// XDG_CONFIG_HOME ▶︎ ~/Library/Preferences/
// XDG_DATA_HOME ▶︎ ~/Library/
// XDG_CACHE_HOME ▶︎ ~/Library/Caches/
// Mapping XDG Base Directory Specification locations for "My App" on Mac OS X could look like this:

func MacConfigHome(user *user.User) string {
	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME == "" {
		return filepath.Join(user.HomeDir, ".config")
	}
	return XDG_CONFIG_HOME
}

func MacCacheHome(user *user.User) string {
	XDG_CACHE_HOME := os.Getenv("XDG_CACHE_HOME")
	if XDG_CACHE_HOME == "" {
		return filepath.Join(user.HomeDir, ".cache")
	}
	return XDG_CACHE_HOME
}

/*
 * Windows standard dirs
 */
// Windows
// const XDG_DATA_HOME = %LOCALAPPDATA%
// $XDG_DATA_DIRS = %APPDATA%
// const XDG_CONFIG_HOME = %LOCALAPPDATA%
// $XDG_CONFIG_DIRS = %APPDATA%
// $XDG_CACHE_HOME = %TEMP%
// $XDG_RUNTIME_DIR = %TEMP%
func WinConfigHome(user *user.User) string {
	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME == "" {
		return filepath.Join(user.HomeDir, ".config")
	}
	return XDG_CONFIG_HOME
}

func WinCacheHome(user *user.User) string {
	XDG_CACHE_HOME := os.Getenv("XDG_CACHE_HOME")
	if XDG_CACHE_HOME == "" {
		return filepath.Join(user.HomeDir, ".cache")
	}
	return XDG_CACHE_HOME
}

// XDG_CONFIG_HOME ▶︎ ~/Library/Preferences/name.often.with.domain.myapp.plist
// XDG_DATA_HOME ▶︎ ~/Library/My App/
// XDG_CACHE_HOME ▶︎ ~/Library/Caches/My App/

func Home() string {
	path := os.Getenv("WK_HOME")
	if path == "" {
		user := currentUser()
		path = ConfigHome(user)
	}
	return path
}

func WkHome() string {
	path := os.Getenv("WK_HOME")
	if path == "" {
		user := currentUser()
		path = ConfigHome(user)
	}
	return path
}

func Projects() string {
	path := os.Getenv("WK_PROJECTS")
	if path == "" {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(user.HomeDir, "projects")
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
func ConfigHome(user *user.User) string {
	switch os := runtime.GOOS; os {
	case "linux":
		return filepath.Join(MacConfigHome(user), "wk")
	case "darwin":
		return filepath.Join(XdgConfigHome(user), "wk")
	case "windows":
		return filepath.Join(WinConfigHome(user), "wk")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		// fmt.Printf("%s.", os)
		return ""
	}

}

func ConfigFilename(filename string) string {
	return ""
}
