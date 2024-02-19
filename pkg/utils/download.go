package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cheggaaa/pb/v3"
)

func Download(url string, target string) (string, error) {
	// Extraer el nombre del archivo
	filePath := Resolve(target, filepath.Base(url))

	// Verificar si el archivo ya existe
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	// Asegurarse de que el directorio de destino exista
	err := os.MkdirAll(target, 0755)
	if err != nil {
		return "", err
	}

	// Obtener el tamaño del archivo para la barra de progreso
	resp, err := http.Head(url)
	if err != nil {
		return "", err
	}
	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return "", err
	}

	// Crear una nueva barra de progreso
	bar := pb.StartNew(size)

	// Descargar el archivo
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Crear el archivo de destino
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Crear un ProxyReader para actualizar la barra de progreso
	reader := bar.NewProxyReader(res.Body)

	// Copiar los datos del archivo .tar al archivo de destino
	if _, err := io.Copy(out, reader); err != nil {
		return "", err
	}

	// Finalizar la barra de progreso
	bar.Finish()

	return filePath, nil
}
