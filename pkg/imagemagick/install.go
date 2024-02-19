package imagemagick

import (
	"fmt"
	"os"
	"runtime"

	"github.com/joalvm/processor-medias/pkg/utils"
)

func Install() {
	existsImageMagick, _ := os.Stat(ImageMagickBinPath)

	if existsImageMagick != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		installWindows()
	default:
		panic("No se puede instalar ImageMagick en este sistema operativo")
	}
}

func installWindows() {
	url := "https://imagemagick.org/archive/binaries/ImageMagick-7.1.1-28-portable-Q16-HDRI-x64.zip"
	fmt.Printf("Descargando ImageMagick desde: %s\n", url)

	imagemagickTempPath, _ := utils.Download(url, os.TempDir())

	fmt.Printf("Descomprimiendo ImageMagick: %s\n", imagemagickTempPath)

	err := utils.UnZip(imagemagickTempPath, os.TempDir(), "imagemagick")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	imagemagickVersion, _ := ImageMagickVersion()

	fmt.Printf("ImageMagick: %s ✔️\n", imagemagickVersion)

	// Eliminar el archivo .zip
	os.Remove(imagemagickTempPath)
}
