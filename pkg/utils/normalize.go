package utils

import "path/filepath"

// Normaliza una ruta de archivo para que sea compatible con el sistema operativo
func Normalize(path string) string {
	return filepath.ToSlash(filepath.FromSlash(path))
}
