package utils

import "github.com/joalvm/processor-medias/pkg/enums"

func GetOrientation(w int, h int) enums.Orientation {
	if w > h {
		return enums.LANDSCAPE
	}

	if w == h {
		return enums.SQUARE
	}

	return enums.PORTRAIT
}
