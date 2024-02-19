package utils

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

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

func UnZip(src string, dest string, mainDirName string) error {
	r, err := zip.OpenReader(src)
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

	for _, f := range r.File {
		if err := processFile(f, dirNamePath, rootDir); err != nil {
			return err
		}
	}
	return nil
}
