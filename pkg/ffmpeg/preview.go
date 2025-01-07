package ffmpeg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

func (f *Ffmpeg) Preview() (string, error) {
	tempDir, err := utils.TempDir()
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDir) // Limpiar después

	fmt.Println("preview file:", f.input)

	// Reemplazar el .mp4 por _preview.mp4
	output := strings.ReplaceAll(f.output, ".mp4", "_preview.mp4")
	clips, err := f.makeClips(tempDir)
	if err != nil {
		return "", err
	}

	return f.concatClips(clips, output)
}

func (f *Ffmpeg) concatClips(clips []string, output string) (string, error) {
	inputs := []string{}
	filters := []string{}

	for _, clip := range clips {
		inputs = append(inputs, "-i", clip)
	}

	for i := 0; i < len(clips); i++ {
		filters = append(filters, fmt.Sprintf("[%d:v:0]", i))
	}

	command := []string{
		"-hide_banner",
		"-loglevel", "error",
		"-v", "error",
	}

	command = append(command, inputs...)

	command = append(
		command,
		[]string{
			"-filter_complex", fmt.Sprintf(
				"%sconcat=n=%d:v=1[outv];[outv]scale=-1:240,pad=ceil(iw/2)*2:ceil(ih/2)*2[outv]",
				strings.Join(filters, ""),
				len(clips),
			),
			"-map", "[outv]",
			"-c:v", "libx264",
			"-crf", "18",
			"-pix_fmt", "yuv420p",
			"-an",
			output,
		}...,
	)

	cmd := exec.Command(FfmpegBinPath, command...)
	cmd.Stderr = os.Stderr

	fmt.Println("command generate Preview:", cmd.String())

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running ffmpeg: %v", err)
	}

	// return concatOutput
	return output, nil
}

func (f *Ffmpeg) makeClips(directory string) ([]string, error) {
	duration := f.Probe.Duration()
	clips := []string{}

	command := []string{
		"-hide_banner",
		"-loglevel", "error",
		"-v", "error",
		"-i", f.input,
	}

	for i, perc := range f.previewMoments {
		moment := int(duration * perc)
		output := filepath.Join(directory, fmt.Sprintf("clip_%d.mp4", i))
		command = append(command, []string{
			"-ss", strconv.Itoa(moment),
			"-t", "1",
			"-c:v", "h264_nvenc",
			"-qp", "18",
			"-vf", "scale=trunc(iw/2)*2:trunc(ih/2)*2",
			"-pix_fmt", "yuv420p",
			"-an",
			output,
		}...)

		clips = append(clips, output)
	}

	cmd := exec.Command(FfmpegBinPath, command...)
	cmd.Stderr = os.Stderr

	fmt.Println("command Generate Clips:", cmd.String())

	err := cmd.Run()
	if err != nil {
		return clips, fmt.Errorf("error running ffmpeg: %v", err)
	}

	return clips, nil
}
