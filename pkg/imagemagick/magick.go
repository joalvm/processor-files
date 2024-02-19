package imagemagick

import (
	"os"
	"os/exec"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var ImageMagickBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "imagemagick/magick"))

// Funcion para ejecutar el comando magick
func Magick(input string, output string, options ...string) error {
	args := []string{input}

	args = append(args, options...)

	if output != "" {
		args = append(args, output)
	}

	cmd := exec.Command(ImageMagickBinPath, args...)

	return cmd.Run()
}

func ImageMagickVersion() (string, error) {
	output, err := exec.Command(ImageMagickBinPath, "--version").Output()

	if err != nil {
		return "", err
	}

	// Dividir la primera línea en saltos de linea y luego en espacios
	parts := strings.Split(strings.Split(string(output), "\n")[0], " ")

	return parts[2], nil
}
