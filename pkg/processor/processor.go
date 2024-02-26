package processor

import (
	"github.com/joalvm/processor-medias/pkg/models"
)

type Processor struct {
	sourceDir      string
	destinationDir string
	ffmpegUrl      string
	imagemagickUrl string
}

var (
	DirectoryMap     = make(map[string][]*models.File)
	DirectoryIndexes = make(map[string]int)
	GlobalIndex      = 1
)

func New(options ...func(*Processor)) *Processor {
	proc := &Processor{}
	for _, o := range options {
		o(proc)
	}
	return proc
}
