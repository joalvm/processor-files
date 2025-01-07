package utils

var mimes = map[string]string{
	// Images
	"JPEG": "image/jpeg",
	"PNG":  "image/png",
	"GIF":  "image/gif",
	"TIFF": "image/tiff",
	"BMP":  "image/bmp",
	"WEBP": "image/webp",
	"SVG":  "image/svg+xml",
	"HEIC": "image/heic",
	"HEIF": "image/heif",
	"AVIF": "image/avif",
	"ICO":  "image/x-icon",
	// Videos
	"MP4":  "video/mp4",
	"WEBM": "video/webm",
	"OGG":  "video/ogg",
	"AVI":  "video/x-msvideo",
	"FLV":  "video/x-flv",
	"MOV":  "video/quicktime",
	"MKV":  "video/x-matroska",
	"3GP":  "video/3gpp",
	"3G2":  "video/3gpp2",
	"WMV":  "video/x-ms-wmv",
	"TS":   "video/mp2t",
	"MPEG": "video/mpeg",
	"MPG":  "video/mpeg",
	"VOB":  "video/mpeg",
	"RM":   "video/vnd.rn-realvideo",
	"RMVB": "video/vnd.rn-realvideo",
	"SWF":  "application/x-shockwave-flash",
	"ASF":  "video/x-ms-asf",
	"AMV":  "video/x-amv",
	"MTS":  "video/mp2t",
	"M2TS": "video/mp2t",
	"QT":   "video/quicktime",
}

func MimeType(mimeType string) string {
	if value, ok := mimes[mimeType]; ok {
		return value
	}

	return ""
}
