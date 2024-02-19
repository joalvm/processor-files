package utils

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func UnZip(path string, dest string, mainDirName string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()

	// Crear el directorio principal
	dirNamePath, err := makeDirectory(dest, mainDirName)
	if err != nil {
		return err
	}

	var rootDir string

	size, _ := countZipFiles(path)

	bar := ProgressBar(size)

	bar.Start()

	for _, f := range r.File {
		bar.Set("filename", f.Name)
		if err := processFile(f, dirNamePath, rootDir); err != nil {
			return err
		}

		bar.Increment()
	}

	bar.Set("filename", "")
	bar.Finish()

	return nil
}

func countZipFiles(path string) (int, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	var count int = 0

	for range r.File {
		count++
	}

	return count, nil
}

func makeDirectory(dest string, mainDirName string) (string, error) {
	dirNamePath := Resolve(dest, mainDirName)

	if err := os.MkdirAll(dirNamePath, os.ModePerm); err != nil {
		return "", err
	}

	return dirNamePath, nil
}

func processFile(f *zip.File, dirNamePath string, rootDir string) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// Obtener el nombre del directorio raíz del archivo .zip
	if rootDir == "" {
		rootDir = strings.Split(f.Name, "/")[0]
	}

	// Si es el directorio raíz, omitir
	if f.Name == rootDir {
		return nil
	}

	// Si es un directorio
	if f.FileInfo().IsDir() {
		// Crear el directorio, omitiendo el nombre del directorio raíz
		newDir := strings.TrimPrefix(f.Name, Normalize(rootDir+"/"))
		if err := os.MkdirAll(Resolve(dirNamePath, newDir), os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	// Crear el archivo, omitiendo el nombre del directorio raíz
	newDir := strings.TrimPrefix(f.Name, rootDir+"/")

	f2, err := os.OpenFile(Resolve(dirNamePath, newDir), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer f2.Close()

	// Copiar los datos al archivo
	if _, err := io.Copy(f2, rc); err != nil {
		return err
	}
	return nil
}
