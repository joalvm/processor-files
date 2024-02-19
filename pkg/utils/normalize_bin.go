package utils

import "runtime"

// Normaliza el nombre de un binario para que sea compatible con el sistema operativo
func NormalizeBin(path string) string {
	os := runtime.GOOS

	if os == "windows" {
		return path + ".exe"
	}

	return path
}
