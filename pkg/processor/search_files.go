package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

/*
SearchFiles is a function that searches for files in a given directory and returns a map of directories and their corresponding files.

Parameters:
- sourceDir (string): The directory to search for files in.

Returns:
- map[string][]string: A map where the keys are directory paths and the values are slices of file names found in each directory.
- error: An error if any occurred during the file search process.

Example Usage:

	directoryMap, err := SearchFiles("/path/to/directory")

	if err != nil {
	    // handle error
	}

// use directoryMap to access the files found in each directory
*/
func SearchFiles(sourceDir string) ([]string, error) {
	var files []string

	// Si el directorio no existe, retornar un error
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory does not exist")
	}

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in directory")
	}

	sort.Strings(files)

	return files, nil
}
