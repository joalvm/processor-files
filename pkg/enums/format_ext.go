package enums

import "encoding/json"

type FormatExt int

const (
	JPEG FormatExt = iota
	WEBP
	GIF
	OGG
	WEBM
	MP4
)

func (fe FormatExt) String() string {
	switch fe {
	case WEBP:
		return "webp"
	case OGG:
		return "ogg"
	case GIF:
		return "gif"
	case WEBM:
		return "webm"
	case MP4:
		return "mp4"
	default:
		return "jpeg"
	}
}

func (fe FormatExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(fe.String())
}
