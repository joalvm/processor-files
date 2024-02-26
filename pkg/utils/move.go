package utils

import (
	"os"
	"path/filepath"
)

func Move(sourceDir string, destinationDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			err := os.MkdirAll(filepath.Join(destinationDir, path[len(sourceDir):]), os.ModePerm)
			if err != nil {
				return err
			}

			return nil
		}

		err = os.Rename(path, filepath.Join(destinationDir, path[len(sourceDir):]))
		if err != nil {
			return err
		}

		return nil
	})
}
