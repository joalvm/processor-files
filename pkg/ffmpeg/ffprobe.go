package ffmpeg

import (
	"os"
	"os/exec"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var FfprobeBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "ffmpeg/bin/ffprobe"))

// Funcion para ejecutar el comando ffprobe
func Ffprobe(input string, options ...string) (string, error) {
	args := []string{"-hide_banner", "-v", "error"}

	args = append(args, options...)

	if input != "" {
		args = append(args, input)
	}

	cmd := exec.Command(FfprobeBinPath, args...)

	output, err := cmd.Output()

	return string(output), err
}

func FfprobeVersion() (string, error) {
	output, err := Ffprobe("", "-version")

	if err != nil {
		return "", err
	}

	// Dividir la primera línea en saltos de linea y luego en espacios
	parts := strings.Split(strings.Split(output, "\n")[0], " ")

	return parts[2], nil
}
