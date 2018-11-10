package fs

import (
	"os"
)

// Exists checks wether or not a path exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsFile checks wether or not a path is a file
func IsFile(path string) bool {
	if !Exists(path) {
		return false
	}
	stats, err := os.Stat(path)
	if err != nil {
		// TODO: Handle erro
	}
	return stats.Mode().IsRegular()
}

// IsDir checks wether or not a path is a directory
func IsDir(path string) bool {
	if !Exists(path) {
		return false
	}
	stats, err := os.Stat(path)
	if err != nil {
		// TODO: Handle erro
	}
	return stats.Mode().IsDir()
}
