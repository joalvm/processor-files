package ffmpeg

type VideoInfo struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

type Format struct {
	Filename       string     `json:"filename"`
	NbStreams      int64      `json:"nb_streams"`
	NbPrograms     int64      `json:"nb_programs"`
	NbStreamGroups int64      `json:"nb_stream_groups"`
	FormatName     string     `json:"format_name"`
	FormatLongName string     `json:"format_long_name"`
	StartTime      string     `json:"start_time"`
	Duration       string     `json:"duration"`
	Size           string     `json:"size"`
	BitRate        string     `json:"bit_rate"`
	ProbeScore     int64      `json:"probe_score"`
	Tags           FormatTags `json:"tags"`
}

type FormatTags struct {
	MajorBrand        string `json:"major_brand"`
	MinorVersion      string `json:"minor_version"`
	CompatibleBrands  string `json:"compatible_brands"`
	CreationTime      string `json:"creation_time"`
	COMAndroidVersion string `json:"com.android.version,omitempty"`
}

type Stream struct {
	Index            int64            `json:"index"`
	CodecName        string           `json:"codec_name"`
	CodecLongName    string           `json:"codec_long_name"`
	Profile          string           `json:"profile"`
	CodecType        string           `json:"codec_type"`
	CodecTagString   string           `json:"codec_tag_string"`
	CodecTag         string           `json:"codec_tag"`
	Width            int              `json:"width,omitempty"`
	Height           int              `json:"height,omitempty"`
	CodedWidth       int              `json:"coded_width,omitempty"`
	CodedHeight      int              `json:"coded_height,omitempty"`
	ClosedCaptions   int64            `json:"closed_captions,omitempty"`
	FilmGrain        int64            `json:"film_grain,omitempty"`
	HasBFrames       int64            `json:"has_b_frames,omitempty"`
	PixFmt           string           `json:"pix_fmt,omitempty"`
	Level            int64            `json:"level,omitempty"`
	ChromaLocation   string           `json:"chroma_location,omitempty"`
	FieldOrder       string           `json:"field_order,omitempty"`
	Refs             int64            `json:"refs,omitempty"`
	IsAVC            string           `json:"is_avc,omitempty"`
	NalLengthSize    string           `json:"nal_length_size,omitempty"`
	ID               string           `json:"id"`
	RFrameRate       string           `json:"r_frame_rate"`
	AvgFrameRate     string           `json:"avg_frame_rate"`
	TimeBase         string           `json:"time_base"`
	StartPts         int64            `json:"start_pts"`
	StartTime        string           `json:"start_time"`
	DurationTs       int64            `json:"duration_ts"`
	Duration         string           `json:"duration"`
	BitRate          string           `json:"bit_rate"`
	BitsPerRawSample string           `json:"bits_per_raw_sample,omitempty"`
	NbFrames         string           `json:"nb_frames"`
	ExtradataSize    int64            `json:"extradata_size"`
	Disposition      map[string]int64 `json:"disposition"`
	Tags             StreamTags       `json:"tags"`
	SideDataList     []SideDataList   `json:"side_data_list,omitempty"`
	SampleFmt        string           `json:"sample_fmt,omitempty"`
	SampleRate       string           `json:"sample_rate,omitempty"`
	Channels         int64            `json:"channels,omitempty"`
	ChannelLayout    string           `json:"channel_layout,omitempty"`
	BitsPerSample    int64            `json:"bits_per_sample,omitempty"`
	InitialPadding   int64            `json:"initial_padding,omitempty"`
}

type SideDataList struct {
	SideDataType  string `json:"side_data_type"`
	Displaymatrix string `json:"displaymatrix"`
	Rotation      int64  `json:"rotation"`
}

type StreamTags struct {
	CreationTime string `json:"creation_time"`
	Language     string `json:"language"`
	HandlerName  string `json:"handler_name"`
	VendorID     string `json:"vendor_id"`
}
