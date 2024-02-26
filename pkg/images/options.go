package images

import (
	"os"

	"github.com/joalvm/processor-medias/pkg/enums"
	"github.com/joalvm/processor-medias/pkg/models"
)

func WithModel(model *models.File) func(*Image) {
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

func WithFormats(formats []enums.FormatExt) func(*Image) {
	return func(i *Image) {
		i.formats = formats
	}
}

func WithThumbSizes(sizes struct{ Md, Sm, Xs int }) func(*Image) {
	return func(i *Image) {
		i.thumbSizes = sizes
	}
}
