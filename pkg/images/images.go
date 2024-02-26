package images

import (
	"fmt"
	"os"
	"strings"

	"github.com/joalvm/processor-medias/pkg/enums"
	"github.com/joalvm/processor-medias/pkg/imagemagick"
	"github.com/joalvm/processor-medias/pkg/models"
	"github.com/joalvm/processor-medias/pkg/utils"
)

type Image struct {
	model          *models.File
	file           *os.File
	directory      string
	destinationDir string
	formats        []enums.FormatExt
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
	w, h, err := imagemagick.NewWithInput(i.file.Name()).Dimensions()
	if err != nil {
		return err
	}

	i.model.Width = w
	i.model.Height = h
	i.model.Orientation = utils.GetOrientation(w, h)
	i.model.AspectRatio = utils.GetAspectRatio(w, h)
	i.model.Preview = nil

	err = i.makeFolder()
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

	err = i.makeFormats()
	if err != nil {
		return err
	}

	return nil
}

func (i *Image) processGif() error {
	i.model.Type = enums.ANIMATED

	return nil
}

func (i *Image) makeFormats() error {
	for _, format := range i.formats {
		name := utils.Resolve(i.getFolder(), fmt.Sprintf("%s.%s", i.model.Name, strings.ToLower(format.String())))

		err := imagemagick.NewWithInput(i.file.Name()).SetOutput(name).Convert(80)
		if err != nil {
			return err
		}

		w, h, err := imagemagick.NewWithInput(name).Dimensions()
		if err != nil {
			return err
		}

		info, err := os.Stat(name)
		if err != nil {
			return err
		}

		i.model.Formats = append(i.model.Formats, models.Format{
			Type:     enums.IMAGE,
			Ext:      format,
			MimeType: fmt.Sprintf("image/%s", strings.ToLower(format.String())),
			Size:     info.Size(),
			Width:    w,
			Height:   h,
		})
	}

	return nil
}

func (i *Image) makeThumbnails() error {
	media, err := i.makeThumbnail(i.thumbSizes.Md, "md")
	if err != nil {
		return err
	}

	i.model.Thumbnails.Md = media

	media, err = i.makeThumbnail(i.thumbSizes.Sm, "sm")
	if err != nil {
		return err
	}

	i.model.Thumbnails.Md = media

	media, err = i.makeThumbnail(i.thumbSizes.Xs, "xs")
	if err != nil {
		return err
	}

	i.model.Thumbnails.Xs = media

	return nil
}

func (i *Image) makeThumbnail(size int, suffix string) (models.Media, error) {
	name := i.thumbName(suffix)

	err := imagemagick.NewWithInput(i.file.Name()).SetOutput(name).Resize(size)
	if err != nil {
		return models.Media{}, err
	}

	w, h, err := imagemagick.NewWithInput(name).Dimensions()
	if err != nil {
		return models.Media{}, err
	}

	info, err := os.Stat(name)
	if err != nil {
		return models.Media{}, err
	}

	return models.Media{
		Width:  w,
		Height: h,
		Size:   info.Size(),
		Path:   name[len(i.getFolder()+string(os.PathSeparator)):],
	}, nil
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
			i.model.Name,
			strings.ToLower(suffix),
			strings.ToLower(ext),
		),
	)
}

func (i *Image) isGif() bool {
	return i.model.MimeType == "image/gif"
}
