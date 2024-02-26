package imagemagick

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var ImageMagickBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "imagemagick/magick"))

type ImageMagick struct {
	input  string
	output string
}

// Funcion para ejecutar el comando magick
func New(options ...func(*ImageMagick)) *ImageMagick {
	proc := &ImageMagick{}
	for _, o := range options {
		o(proc)
	}
	return proc
}

func NewWithInput(input string) *ImageMagick {
	return New(WithInput(input))
}

func (im *ImageMagick) SetOutput(output string) *ImageMagick {
	im.output = output
	return im
}

func WithInput(input string) func(*ImageMagick) {
	return func(i *ImageMagick) {
		i.input = input
	}
}

func WithOutput(output string) func(*ImageMagick) {
	return func(i *ImageMagick) {
		i.output = output
	}
}

func (im *ImageMagick) Version() (string, error) {
	output, err := im.out("--version")
	if err != nil {
		return "", err
	}

	// Dividir la primera línea en saltos de linea y luego en espacios
	parts := strings.Split(strings.Split(string(output), "\n")[0], " ")

	return parts[2], nil
}

func (im *ImageMagick) run(args ...string) error {
	return im.cmd(args...).Run()
}

func (im *ImageMagick) out(args ...string) ([]byte, error) {
	return im.cmd(args...).Output()
}

func (im *ImageMagick) cmd(args ...string) *exec.Cmd {
	return exec.Command(ImageMagickBinPath, args...)
}

func (im *ImageMagick) validateInAndOutput() error {
	if im.input == "" {
		return fmt.Errorf("input is required")
	}

	if im.output == "" {
		return fmt.Errorf("output is required")
	}

	return nil
}
