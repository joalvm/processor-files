package ffmpeg

import (
	"os"
	"os/exec"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var (
	FfmpegBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "ffmpeg/bin/ffmpeg"))
)

type Ffmpeg struct {
	Probe          *Ffprobe
	input          string
	output         string
	previewMoments []float64
}

// Funcion para ejecutar el comando ffmpeg
func NewFfmpeg() *Ffmpeg {
	return &Ffmpeg{
		Probe:          NewFfprobe(),
		input:          "",
		output:         "",
		previewMoments: []float64{0.10, 0.30, 0.50, 0.70, 0.90},
	}
}

func (f *Ffmpeg) Input(input string) *Ffmpeg {
	f.input = input
	f.Probe = f.Probe.Input(input)

	return f
}

func (f *Ffmpeg) Output(output string) *Ffmpeg {
	f.output = output
	return f
}

func (f *Ffmpeg) run(args ...string) error {
	return f.cmd(args...).Run()
}

func (f *Ffmpeg) out(args ...string) ([]byte, error) {
	return f.cmd(args...).Output()
}

func (f *Ffmpeg) cmd(args ...string) *exec.Cmd {
	arguments := []string{"-hide_banner", "-loglevel", "error", "-v", "error"}

	arguments = append(arguments, args...)

	cmd := exec.Command(FfmpegBinPath, arguments...)

	cmd.Stderr = os.Stderr

	return cmd
}
