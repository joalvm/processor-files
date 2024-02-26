package models

import (
	"github.com/joalvm/processor-medias/pkg/enums"
)

type File struct {
	Code        string            `json:"code"`
	Name        string            `json:"Name"`
	RealName    string            `json:"real_name"`
	Type        enums.FileType    `json:"type"`
	Index       int               `json:"index"`
	DirIndex    int               `json:"dir_index"`
	Width       int               `json:"width"`
	Height      int               `json:"height"`
	AspectRatio AspectRatio       `json:"aspect_ratio"`
	Size        int64             `json:"size"`
	Orientation enums.Orientation `json:"orientation"`
	Duration    int               `json:"duration,omitempty"`
	Bitrate     int               `json:"bitrate,omitempty"`
	MimeType    string            `json:"mime_type"`
	Codec       string            `json:"codec,omitempty"`
	AudioCodec  string            `json:"audio_codec,omitempty"`
	Resolution  string            `json:"resolution,omitempty"`
	Fps         string            `json:"fps,omitempty"`
	IsMuted     bool              `json:"is_muted,omitempty"`
	Preview     *Media            `json:"preview,omitempty"`
	Thumbnails  Thumbnails        `json:"thumbnails"`
	Formats     []Format          `json:"formats"`
}
