package utils

import (
	"path/filepath"
)

func Resolve(paths ...string) string {
	joinedPath := filepath.Join(paths...)
	absPath, err := filepath.Abs(joinedPath)

	if err != nil {
		return ""
	}

	return absPath
}
