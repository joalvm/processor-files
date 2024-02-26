package models

import "github.com/joalvm/processor-medias/pkg/enums"

type Format struct {
	Type       enums.FileType  `json:"type"`
	Ext        enums.FormatExt `json:"ext"`
	MimeType   string          `json:"mime_type"`
	Size       int64           `json:"size"`
	Width      int             `json:"width"`
	Height     int             `json:"height"`
	Bitrate    int             `json:"bitrate,omitempty"`
	Codec      string          `json:"codec,omitempty"`
	IsMuted    bool            `json:"is_muted,omitempty"`
	Resolution string          `json:"resolution,omitempty"`
	Fps        int             `json:"fps,omitempty"`
}
