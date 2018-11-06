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
	stats, err := os.Stat(path)
	if err != nil {
		// TODO: Handle erro
	}
	return stats.Mode().IsRegular()
}

// IsFile checks wether or not a path is a directory
func IsDir(path string) bool {
	stats, err := os.Stat(path)
	if err != nil {
		// TODO: Handle erro
	}
	return stats.Mode().IsDir()
}
