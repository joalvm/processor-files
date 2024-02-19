package utils

import (
	"io"
	"os"
	"strings"

	"github.com/saracen/go7z"
	"github.com/saracen/go7z/headers"
)

func Un7z(path string, target string, mainDirName string) (string, error) {
	sz, err := go7z.OpenReader(path)
	if err != nil {
		return "", err
	}
	defer sz.Close()

	// Crear el directorio principal
	dirNamePath, err := createDirectory(target, mainDirName)
	if err != nil {
		return "", err
	}

	var rootDir string

	for {
		hdr, err := sz.Next()
		if err == io.EOF {
			break // Fin del archivo
		}
		if err != nil {
			return "", err
		}

		// Obtener el nombre del directorio raíz del archivo .7z
		if rootDir == "" {
			rootDir = strings.Split(hdr.Name, "/")[0]
		}

		processHeader(sz, dirNamePath, hdr, rootDir)
	}

	return dirNamePath, nil
}

func createDirectory(target string, mainDirName string) (string, error) {
	dirNamePath := Resolve(target, mainDirName)

	if err := os.MkdirAll(dirNamePath, os.ModePerm); err != nil {
		return "", err
	}

	return dirNamePath, nil
}

func processHeader(sz *go7z.ReadCloser, dirNamePath string, hdr *headers.FileInfo, rootDir string) error {
	// Si es el directorio raíz, omitir
	if hdr.Name == rootDir {
		return nil
	}

	// Si es un directorio
	if hdr.IsEmptyStream && !hdr.IsEmptyFile {
		// Crear el directorio, omitiendo el nombre del directorio raíz
		newDir := strings.TrimPrefix(hdr.Name, Normalize(rootDir+"/"))

		if err := os.MkdirAll(Resolve(dirNamePath, newDir), os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	// Crear el archivo, omitiendo el nombre del directorio raíz
	newDir := strings.TrimPrefix(hdr.Name, rootDir+"/")
	f, err := os.Create(Resolve(dirNamePath, newDir))
	if err != nil {
		return err
	}
	defer f.Close()

	// Copiar los datos al archivo
	if _, err := io.Copy(f, sz); err != nil {
		return err
	}

	return nil
}
