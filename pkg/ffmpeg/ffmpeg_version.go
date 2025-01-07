package ffmpeg

import "strings"

func (f *Ffmpeg) Version() (string, error) {
	output, err := f.out("-version")
	if err != nil {
		return "", err
	}

	// Dividir la primera línea en saltos de linea y luego en espacios
	parts := strings.Split(strings.Split(string(output), "\n")[0], " ")

	return parts[2], nil
}
