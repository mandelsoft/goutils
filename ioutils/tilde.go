package ioutils

import (
	"fmt"
	"os"
	"strings"
)

// ResolvePath handles the ~ notation for the home directory.
func ResolvePath(path string) (string, error) {
	if strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		home, err := os.UserHomeDir()
		if home == "" || err != nil {
			return path, fmt.Errorf("HOME not set")
		}
		path = home + path[1:]
	}
	return path, nil
}
