package processor

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/database"
	"github.com/joalvm/processor-medias/pkg/models"
	"github.com/joalvm/processor-medias/pkg/utils"
	"gorm.io/gorm"
)

type Processor struct {
	DirectoryMap       map[string][]*models.Media
	FilesIndexer       map[string]int
	DirectoriesIndexer map[string]int
	GlobalIndex        int
	sourceDir          string
	destinationDir     string
	ffmpegUrl          string
	imagemagickUrl     string
	db                 *gorm.DB
	directoryService   *database.DirectoryService
}

func New(options ...func(*Processor)) *Processor {
	proc := &Processor{}
	for _, o := range options {
		o(proc)
	}

	proc.DirectoryMap = make(map[string][]*models.Media)
	proc.FilesIndexer = make(map[string]int)
	proc.DirectoriesIndexer = make(map[string]int)
	proc.directoryService = database.NewDirectoryService(proc.db)

	return proc
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

func WithDb(db *gorm.DB) func(*Processor) {
	return func(p *Processor) {
		p.db = db
	}
}

func (p *Processor) setDirectoryIndex(directory string) {
	if directory == p.sourceDir {
		return
	}

}
