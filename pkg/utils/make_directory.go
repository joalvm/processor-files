package utils

import "os"

func MakeDirectory(path string) error {
	// Verificar si el directorio existe
	if _, err := os.Stat(path); os.IsExist(err) {
		return nil
	}

	// Crear el directorio
	return os.MkdirAll(path, os.ModePerm)
}
