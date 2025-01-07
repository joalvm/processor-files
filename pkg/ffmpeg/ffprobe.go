package ffmpeg

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/joalvm/processor-medias/pkg/utils"
)

var FfprobeBinPath string = utils.NormalizeBin(utils.Resolve(os.TempDir(), "ffmpeg/bin/ffprobe"))

type Ffprobe struct {
	info  VideoInfo
	input string
}

// Funcion para ejecutar el comando ffprobe
func NewFfprobe() *Ffprobe {
	return &Ffprobe{
		info:  VideoInfo{},
		input: "",
	}
}

func (fp *Ffprobe) Input(input string) *Ffprobe {
	fp.input = input

	info, _ := fp.Info()
	fp.info = info

	return fp
}

func (fp *Ffprobe) Info() (VideoInfo, error) {
	// Verificar si el input existe
	if _, err := os.Stat(fp.input); os.IsNotExist(err) {
		return VideoInfo{}, err
	}

	cmd := exec.Command(
		FfprobeBinPath,
		[]string{
			"-hide_banner", "-loglevel", "error",
			"-v", "error",
			"-show_format",
			"-show_streams",
			"-of", "json",
			fp.input,
		}...,
	)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error al ejecutar el comando ffprobe", err)
		return VideoInfo{}, err
	}

	info := VideoInfo{}
	err = json.Unmarshal(out, &info)
	if err != nil {
		fmt.Println("Error al convertir el json", err)
		return VideoInfo{}, err
	}

	return info, nil
}

func (fp *Ffprobe) Version() (string, error) {
	out, err := exec.Command(FfprobeBinPath, []string{"-hide_banner", "-loglevel", "error", "-version"}...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func (fp *Ffprobe) StreamVideo() Stream {
	for _, stream := range fp.info.Streams {
		if stream.CodecType == "video" {
			return stream
		}
	}

	return Stream{}
}

func (fp *Ffprobe) StreamAudio() Stream {
	for _, stream := range fp.info.Streams {
		if stream.CodecType == "audio" {
			return stream
		}
	}

	return Stream{}
}

func (fp *Ffprobe) Duration() float64 {
	val, err := strconv.ParseFloat(fp.info.Format.Duration, 64)
	if err != nil {
		return 0
	}

	return val
}

func (fp *Ffprobe) Fps() float64 {
	stream := fp.StreamVideo()
	out := stream.AvgFrameRate

	split := strings.Split(string(out), "/")
	numerator, _ := strconv.Atoi(split[0])
	denominator, _ := strconv.Atoi(strings.TrimSpace(split[1]))

	// If the average frame rate is not available, try to get the frame rate
	if numerator == 0 {
		out = stream.RFrameRate

		split = strings.Split(string(out), "/")
		numerator, _ = strconv.Atoi(split[0])
		denominator, _ = strconv.Atoi(strings.TrimSpace(split[1]))
	}

	return float64(numerator) / float64(denominator)
}

func (fp *Ffprobe) Bitrate() (int64, error) {
	out := fp.info.Format.BitRate

	return strconv.ParseInt(string(out), 10, 64)
}

func (fp *Ffprobe) Dimesions() (int, int) {
	stream := fp.StreamVideo()

	return stream.Width, stream.Height
}
