package images

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/models"
	"gorm.io/gorm"
)

func WithModel(model *models.Media) func(*Image) {
	return func(i *Image) {
		i.model = model
	}
}

func WithFile(file *os.File) func(*Image) {
	return func(i *Image) {
		i.file = file
	}
}

func WithDirectory(directory string) func(*Image) {
	return func(i *Image) {
		i.directory = directory
	}
}

func WithDestinationDir(destinationDir string) func(*Image) {
	return func(i *Image) {
		i.destinationDir = destinationDir
	}
}

func WithThumbSizes(sizes struct{ Md, Sm, Xs int }) func(*Image) {
	return func(i *Image) {
		i.thumbSizes = sizes
	}
}

func WithDb(db *gorm.DB) func(*Image) {
	return func(i *Image) {
		i.db = db
	}
}
