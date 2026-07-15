package config

import "7DL/ffmpeg"

var VideoCodecs = []ffmpeg.VideoCodec{
	"libx265",
	"hevc_nvenc",
	"hevc_amf",
	"hevc_vaapi",
	"libx264",
	"h264_nvenc",
	"h264_amf",
	"h264_vaapi",
	"libaom-av1",
	"libsvtav1",
}

var AudioCodecs = []ffmpeg.AudioCodec{
	"aac",
	"flac",
	"libmp3lame",
	"libopus",
	"ac3",
	"eac3",
}

var noaudio ffmpeg.AudioCodec = "anull"
