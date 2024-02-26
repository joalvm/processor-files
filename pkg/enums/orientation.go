package enums

import "encoding/json"

type Orientation int

const (
	LANDSCAPE Orientation = iota
	PORTRAIT
	SQUARE
)

func (o Orientation) String() string {
	switch o {
	case LANDSCAPE:
		return "landscape"
	case PORTRAIT:
		return "portrait"
	default:
		return "square"
	}
}

func (o Orientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())
}
