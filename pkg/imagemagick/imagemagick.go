package imagemagick

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var ImageMagickBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "imagemagick/magick"))

type ImageInfo struct {
	Path   string
	Width  int
	Height int
	Size   int64
	Mime   string
	Ext    string
}

type ImageMagick struct {
	Info   ImageInfo
	input  string
	output string
	args   []string
}

// Funcion para ejecutar el comando magick
func NewImageMagick() *ImageMagick {
	return &ImageMagick{
		Info:   ImageInfo{},
		input:  "",
		output: "",
		args:   []string{},
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

func (im *ImageMagick) Input(input string) *ImageMagick {
	im.input = input

	im.loadInfo()

	return im
}

func (im *ImageMagick) Output(output string) *ImageMagick {
	im.output = output

	return im
}

func (im *ImageMagick) Resize(size int) *ImageMagick {
	im.args = append(im.args, "-resize", fmt.Sprintf("%dx%d", size, size))
	return im
}

func (im *ImageMagick) Quality(quality int) *ImageMagick {
	im.args = append(im.args, "-quality", fmt.Sprintf("%d", quality))
	return im
}

func (im *ImageMagick) Save() (ImageInfo, error) {
	if im.input == "" {
		return ImageInfo{}, fmt.Errorf("no se ha especificado una imagen")
	}

	if im.output == "" {
		return ImageInfo{}, fmt.Errorf("no se ha especificado un archivo de salida")
	}

	im.args = append([]string{im.input}, im.args...)

	err := im.run(append(im.args, im.output)...)
	if err != nil {
		return ImageInfo{}, err
	}

	return NewImageMagick().Input(im.output).Info, nil
}

func (im *ImageMagick) loadInfo() {
	if im.input == "" {
		return
	}

	cmd := exec.Command(ImageMagickBinPath, "identify", "-format", "%w|%h|%B|%m|%e", im.input)
	out, err := cmd.Output()
	if err != nil {
		return
	}

	parts := strings.Split(string(out), "|")

	width, _ := strconv.Atoi(parts[0])
	height, _ := strconv.Atoi(parts[1])
	size, _ := strconv.ParseInt(parts[2], 10, 64)

	im.Info = ImageInfo{
		Path:   im.input,
		Width:  width,
		Height: height,
		Size:   size,
		Mime:   utils.MimeType(parts[3]),
		Ext:    "." + parts[4],
	}
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
