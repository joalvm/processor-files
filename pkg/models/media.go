package models

import (
	"time"

	"github.com/joalvm/processor-medias/pkg/enums"
)

type Media struct {
	Id           int            `json:"id"`
	DirectoryId  int            `json:"directory_id"`
	Code         string         `json:"code"`
	RealName     string         `json:"real_name"`
	Type         enums.FileType `json:"type"`
	Index        int            `json:"index"`
	DirIndex     int            `json:"dir_index"`
	Size         int64          `json:"size"`
	MimeType     string         `json:"mime_type"`
	Width        int            `json:"width"`
	Height       int            `json:"height"`
	Duration     float64        `json:"duration"`
	Bitrate      int            `json:"bitrate"`
	Codec        string         `json:"codec"`
	AudioCodec   string         `json:"audio_codec"`
	Resolution   string         `json:"resolution"`
	Fps          float64        `json:"fps"`
	IsMuted      bool           `json:"is_muted"`
	LastModified time.Time      `json:"last_modified"`
}
