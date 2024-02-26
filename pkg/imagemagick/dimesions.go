package imagemagick

import (
	"fmt"
	"strconv"
	"strings"
)

func (im *ImageMagick) Dimensions() (int, int, error) {
	if im.input == "" {
		return 0, 0, fmt.Errorf("no se ha especificado una imagen")
	}

	// Crear el comando
	out, err := im.out("identify", "-format", "%wx%h", im.input)
	if err != nil {
		return 0, 0, err
	}

	// Parsear la salida para obtener el ancho y la altura
	dimensions := strings.Split(string(out), "x")

	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error al convertir el ancho de la imagen a entero: %v", err)
	}

	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return width, 0, fmt.Errorf("error al convertir la altura de la imagen a entero: %v", err)
	}

	return width, height, nil
}
