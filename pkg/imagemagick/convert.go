package imagemagick

import "fmt"

func (im *ImageMagick) Convert(quality int) error {
	if err := im.validateInAndOutput(); err != nil {
		return err
	}

	return im.run(im.input, "-quality", fmt.Sprintf("%d", quality), im.output)
}
