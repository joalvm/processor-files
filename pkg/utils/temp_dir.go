package utils

import (
	"fmt"
	"os"
	"strings"
)

func TempDir() (string, error) {
	dirname := GetCode()

	tempDir, err := os.MkdirTemp("", strings.ReplaceAll(dirname, "-", "_"))
	if err != nil {
		return "", fmt.Errorf("error creating temp dir: %v", err)
	}

	return tempDir, nil
}
