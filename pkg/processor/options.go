package processor

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/utils"
)

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
