package processor

import (
	"encoding/json"
	"os"

	"github.com/joalvm/processor-medias/pkg/ffmpeg"
	"github.com/joalvm/processor-medias/pkg/imagemagick"
	"github.com/joalvm/processor-medias/pkg/utils"
)

type Processor struct {
	sourceDir      string
	destinationDir string
	ffmpegUrl      string
	imagemagickUrl string
}

func New(options ...func(*Processor)) *Processor {
	proc := &Processor{}
	for _, o := range options {
		o(proc)
	}
	return proc
}

func (p *Processor) HandleThirdPartyLibraries() error {
	err := ffmpeg.Install(p.ffmpegUrl)
	if err != nil {
		return err
	}

	err = imagemagick.Install(p.imagemagickUrl)
	if err != nil {
		return err
	}

	return nil
}

func (p *Processor) Process() error {
	err := p.HandleThirdPartyLibraries()
	if err != nil {
		return err
	}

	directoryMap, err := SearchFiles(p.sourceDir)

	if err != nil {
		return err
	}

	// Crear un archivo json con la estructura de directorios y archivos
	jsonData, err := json.Marshal(directoryMap)
	if err != nil {
		return err
	}

	f, err := os.Create(utils.Normalize(p.destinationDir + "new_files.json"))
	if err != nil {
		return err
	}

	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func WithSourceDir(sourceDir string) func(*Processor) {
	return func(p *Processor) {
		if string(sourceDir[len(sourceDir)-1]) != string(os.PathSeparator) {
			sourceDir = sourceDir + "/"
		}

		p.sourceDir = utils.Normalize(sourceDir)
	}
}

func WithDestinationDir(destinationDir string) func(*Processor) {
	return func(p *Processor) {
		if string(destinationDir[len(destinationDir)-1]) != string(os.PathSeparator) {
			destinationDir = destinationDir + "/"
		}
		p.destinationDir = utils.Normalize(destinationDir)
	}
}

func WithFfmpegUrl(ffmpegUrl string) func(*Processor) {
	return func(p *Processor) {
		p.ffmpegUrl = ffmpegUrl
	}
}

func WithImagemagickUrl(imagemagickUrl string) func(*Processor) {
	return func(p *Processor) {
		p.imagemagickUrl = imagemagickUrl
	}
}
