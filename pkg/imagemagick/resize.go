package imagemagick

import "fmt"

func (im *ImageMagick) Resize(size int) error {
	if err := im.validateInAndOutput(); err != nil {
		return err
	}

	return im.run("convert", im.input, "-resize", fmt.Sprintf("%dx%d", size, size), "-quality", "90", im.output)
}
