package ffmpeg

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func (f *Ffmpeg) Screenshot() (string, error) {
	output := strings.Replace(f.output, ".mp4", "", 1) + "_screenshot.jpg"

	time := strconv.Itoa(int(f.Probe.Duration() * 0.10))

	cmd := exec.Command(FfmpegBinPath, "-i", f.input, "-ss", time, "-vframes", "1", output)

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error al momento de capturar el screenshot: %v", err)
	}

	return output, nil
}
