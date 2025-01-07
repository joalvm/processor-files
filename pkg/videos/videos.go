package videos

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/enums"
	"github.com/joalvm/processor-medias/pkg/models"
	"gorm.io/gorm"
)

type Video struct {
	model          *models.Media
	file           *os.File
	directory      string
	destinationDir string
	formats        []enums.FormatExt
	thumbSizes     struct{ Md, Sm, Xs int }
	thumbFormat    enums.FormatExt
	db             *gorm.DB
}

func New(options ...func(*Video)) *Video {
	proc := &Video{}
	for _, o := range options {
		o(proc)
	}
	return proc
}

func WithModel(model *models.Media) func(*Video) {
	return func(v *Video) {
		v.model = model
	}
}

func WithFile(file *os.File) func(*Video) {
	return func(v *Video) {
		v.file = file
	}
}

func WithDirectory(directory string) func(*Video) {
	return func(v *Video) {
		v.directory = directory
	}
}

func WithDestinationDir(destinationDir string) func(*Video) {
	return func(v *Video) {
		v.destinationDir = destinationDir
	}
}

func WithFormats(formats []enums.FormatExt) func(*Video) {
	return func(v *Video) {
		v.formats = formats
	}
}

func WithThumbSizes(sizes struct{ Md, Sm, Xs int }) func(*Video) {
	return func(v *Video) {
		v.thumbSizes = sizes
	}
}

func WithThumbFormat(format enums.FormatExt) func(*Video) {
	return func(v *Video) {
		v.thumbFormat = format
	}
}

func WithDb(db *gorm.DB) func(*Video) {
	return func(v *Video) {
		v.db = db
	}
}
