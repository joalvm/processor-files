package enums

import "encoding/json"

type FileType int

const (
	IMAGE FileType = iota
	VIDEO
	AUDIO
	ANIMATED
)

func (ft FileType) String() string {
	switch ft {
	case VIDEO:
		return "video"
	case AUDIO:
		return "audio"
	case ANIMATED:
		return "animated"
	default:
		return "image"
	}
}

// Definir el método MarshalJSON() para el tipo de dato FileType
func (ft FileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ft.String())
}
