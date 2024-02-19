package ffmpeg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/joalvm/processor-medias/pkg/utils"
)

func Install() {
	existsFfmpeg, _ := os.Stat(FfmpegBinPath)
	existsFfprobe, _ := os.Stat(FfprobeBinPath)

	if existsFfprobe != nil || existsFfmpeg != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		installWindows()
	default:
		panic("No se puede instalar ffmpeg y ffprobe en este sistema operativo")
	}
}

func installWindows() {
	url := "https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-full.7z"
	fmt.Printf("Descargando ffmpeg y ffprobe desde: %s\n", url)

	ffmpegTempPath, _ := utils.Download(url, os.TempDir())

	fmt.Printf("Descomprimiendo ffmpeg: %s\n", ffmpegTempPath)

	_, err := utils.Un7z(ffmpegTempPath, os.TempDir(), "ffmpeg")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	ffmpegVersion, _ := FfmpegVersion()
	ffprobeVersion, _ := FfprobeVersion()

	fmt.Printf("ffmpeg: %s ✔️\n", ffmpegVersion)
	fmt.Printf("ffprobe: %s ✔️\n", ffprobeVersion)

	// Eliminar el archivo .7z
	os.Remove(ffmpegTempPath)
}
