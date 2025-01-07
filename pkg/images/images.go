package images

import (
	"fmt"
	"os"
	"strings"

	"github.com/joalvm/processor-medias/pkg/enums"
	"github.com/joalvm/processor-medias/pkg/imagemagick"
	"github.com/joalvm/processor-medias/pkg/models"
	"github.com/joalvm/processor-medias/pkg/utils"
	"gorm.io/gorm"
)

type Image struct {
	model          *models.Media
	file           *os.File
	db             *gorm.DB
	directory      string
	destinationDir string
	thumbSizes     struct{ Md, Sm, Xs int }
	thumbFormat    enums.FormatExt
}

func New(options ...func(*Image)) *Image {
	proc := &Image{}
	for _, o := range options {
		o(proc)
	}
	return proc
}

func (i *Image) Process() error {
	magick := imagemagick.NewImageMagick().Input(i.file.Name())

	i.model.Width = magick.Info.Width
	i.model.Height = magick.Info.Height
	i.model.MimeType = magick.Info.Mime

	err := i.makeFolder()
	if err != nil {
		return err
	}

	if i.isGif() {
		return i.processGif()
	}

	return i.processImage()
}

func (i *Image) processImage() error {
	i.model.Type = enums.IMAGE

	err := i.makeThumbnails()
	if err != nil {
		return err
	}

	name := utils.Resolve(i.getFolder(), fmt.Sprintf("%s.%s", i.model.Code, strings.ToLower(i.thumbFormat.String())))

	_, err = imagemagick.NewImageMagick().Input(i.file.Name()).Output(name).Quality(80).Save()
	if err != nil {
		return err
	}

	return nil
}

func (i *Image) processGif() error {
	i.model.Type = enums.ANIMATED

	return nil
}

func (i *Image) makeThumbnails() error {
	// Iterar todos los tamaños de miniaturas y generarlos

	for _, size := range []struct {
		size   int
		suffix string
	}{
		{i.thumbSizes.Md, "md"},
		{i.thumbSizes.Sm, "sm"},
		{i.thumbSizes.Xs, "xs"},
	} {
		err := i.makeThumbnail(size.size, size.suffix)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Image) makeThumbnail(size int, suffix string) error {
	name := i.thumbName(suffix)

	_, err := imagemagick.NewImageMagick().Input(i.file.Name()).Output(name).Resize(size).Quality(90).Save()
	if err != nil {
		return err
	}

	return nil
}

func (i *Image) makeFolder() error {
	err := os.MkdirAll(i.getFolder(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (i *Image) getFolder() string {
	return utils.Resolve(i.destinationDir, i.directory, i.model.Code)
}

func (i *Image) thumbName(suffix string) string {
	ext := i.thumbFormat.String()
	if ext == "" {
		ext = enums.JPEG.String()
	}

	return utils.Resolve(
		i.getFolder(),
		fmt.Sprintf(
			"%s_%s.%s",
			i.model.Code,
			strings.ToLower(suffix),
			strings.ToLower(ext),
		),
	)
}

func (i *Image) isGif() bool {
	return i.model.MimeType == "image/gif"
}
