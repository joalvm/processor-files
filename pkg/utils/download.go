package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Download downloads a file from the given URL and saves it to the specified target location.
// It includes a progress bar to track the download progress.
//
// Parameters:
//   - url (string): The URL of the file to be downloaded.
//   - target (string): The target directory where the downloaded file will be saved.
//
// Returns:
//   - string: The path of the downloaded file.
//   - error: An error if any occurred during the download process.
func Download(url string, target string) (string, error) {
	// Resolve the file path
	filePath := Resolve(target, filepath.Base(url))

	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	// Ensure that the target directory exists
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return "", err
	}

	// Get the file size for the progress bar
	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return "", err
	}

	// Create a new progress bar
	bar := ProgressBar(size)
	bar.Start()

	// Download the file
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Create the destination file
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Create a ProxyReader to update the progress bar
	reader := bar.NewProxyReader(res.Body)

	// Copy the .tar file data to the destination file
	if _, err := io.Copy(out, reader); err != nil {
		return "", err
	}

	// Finish the progress bar
	bar.Finish()

	return filePath, nil
}
