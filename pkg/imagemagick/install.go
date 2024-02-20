package imagemagick

import (
	"fmt"
	"os"
	"runtime"

	"github.com/joalvm/processor-medias/pkg/utils"
)

func Install(url string) error {
	existsImageMagick, _ := os.Stat(ImageMagickBinPath)

	if existsImageMagick != nil {
		return nil
	}

	switch runtime.GOOS {
	case "windows":
		return installWindows(url)
	default:
		return fmt.Errorf("OS not supported")
	}
}

func installWindows(url string) error {
	fmt.Printf("Descargando ImageMagick desde: %s\n", url)

	imagemagickTempPath, _ := utils.Download(url, os.TempDir())

	fmt.Printf("Descomprimiendo ImageMagick: %s\n", imagemagickTempPath)

	err := utils.UnZip(imagemagickTempPath, os.TempDir(), "imagemagick")

	if err != nil {
		return err
	}

	imagemagickVersion, _ := ImageMagickVersion()

	fmt.Printf("ImageMagick: %s ✔️\n", imagemagickVersion)

	// Eliminar el archivo .zip
	os.Remove(imagemagickTempPath)

	return nil
}
