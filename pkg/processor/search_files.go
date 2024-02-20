package processor

import (
	"os"
	"path/filepath"
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
func SearchFiles(sourceDir string) (map[string][]string, error) {
	directoryMap := make(map[string][]string)

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			directory := path[len(sourceDir):]
			if directory == "" {
				directory = string(os.PathSeparator)
			}

			if _, exists := directoryMap[directory]; !exists {
				directoryMap[directory] = []string{}
			}

			return nil
		}

		directory := filepath.Dir(path)[len(sourceDir):]

		if directory == "" {
			directory = string(os.PathSeparator)
		}

		directoryMap[directory] = append(directoryMap[directory], filepath.Base(path))

		return nil
	})

	if err != nil {
		return nil, err
	}

	return directoryMap, nil
}
