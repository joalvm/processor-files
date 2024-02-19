package ffmpeg

import (
	"os"
	"os/exec"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var (
	FfmpegBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "ffmpeg/bin/ffmpeg"))
)

// Funcion para ejecutar el comando ffmpeg
func Ffmpeg(inputs []string, output string, options ...string) error {
	args := []string{"-y", "-hide_banner", "-loglevel", "error", "-v", "error"}

	// Agregamos los inputs
	for _, input := range inputs {
		args = append(args, "-i", input)
	}

	args = append(args, options...)

	args = append(args, output)

	cmd := exec.Command(FfmpegBinPath, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func FfmpegVersion() (string, error) {
	output, err := exec.Command(FfmpegBinPath, "-hide_banner", "-version").Output()

	if err != nil {
		return "", err
	}

	// Dividir la primera línea en saltos de linea y luego en espacios
	parts := strings.Split(strings.Split(string(output), "\n")[0], " ")

	return parts[2], nil
}
